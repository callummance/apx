package auth

import (
	"errors"
	"fmt"
	"github.com/callummance/apx-srv/models"
	"github.com/callummance/apx-srv/db"
	"github.com/gin-gonic/gin"
	"time"
)

const defaultAvatar string = ""
const sessionCookieName string = "apx_session"

//Generates a new user struct from data fetched from Facebook
func BuildUser(details FbDetails, c *db.DbConn) models.User {
	userDetails := getUserAdvancedDetails(details.User_id)
	var newUser models.User
	newUser.Id = c.GetUUID()
	newUser.FId = details.User_id
	newUser.Name = userDetails.Name
	newUser.Email = userDetails.Email
	newUser.Avatar = defaultAvatar

	return newUser
}


func AuthHandler(c *gin.Context) {
  //Get the authentication token from the request
  authToken := c.Query("auth_token")
  //Retrieve the database connection
  rdb := db.ReactSession
  
  uid, err := AuthenticateUser(authToken, rdb)
  if err != nil {
    //Token is invalid
    panic("Token error, should probably handle this better...")
  }
  sessionKey := NewSession(uid, rdb)
  c.SetCookie(sessionCookieName, sessionKey.SessionKey, sessionDuration, "/", "apx.twintailsare.moe", false, false)
}



//Authenticates an existing user session and retrieves the
//user ID it was assigned to
func AuthSession(c *gin.Context, rdb *db.DbConn) (string, error) {
  //Get cookie from the gin router
  sessionKey, err := c.Cookie(sessionCookieName)
  if err != nil {
    return "", err
    //Cookie not found, apparently...
  }
  
  //Lookup session in database
  session, found, err := rdb.GetSession(sessionKey)
  if err != nil {
    fmt.Println("couldnt run query")
    return "", err
  }
  if !found {
    fmt.Println("no matching cookies")
    return "", errors.New("No such session")
  } else {
    curTime := time.Now().UTC().Unix()
    expireTime := session.Expires
    if (expireTime < curTime) {
      return "", errors.New("Session expired")
    } else {
      fmt.Println(session.UID)
      return session.UID, nil
    }
  }
}

func AuthenticateUser(token string, rdb *db.DbConn) (string, error) {
  userDetails := GetUserDetails(token)
  if !userDetails.Is_valid {
    return "", errors.New("Token provided is not valid or has expired.")
  }
  
  //Lookup the user in database
  fid := userDetails.User_id
  user, found, err := rdb.GetFBUser(fid)
  if err != nil {
      return "", err
  } else if found {
    //If the user exists, return the user id and nil
    return user.Id, nil
  } else {
    //Otherwise, make a new user with BuildUser, and add it to the database
    newUser := BuildUser(userDetails, rdb)
    rdb.WriteUser(newUser)
    user = &newUser
    fmt.Println("Writing new user")
  }
  return user.Id, nil
}
