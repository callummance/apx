package auth

import (
	"errors"
	"fmt"
	"github.com/callummance/apx-srv/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const defaultAvatar string = ""
const sessionCookieName string = "apx_session"

//Generates a new user struct from data fetched from Facebook
func BuildUser(details FbDetails) models.User {
	userDetails := getUserAdvancedDetails(details.User_id)
	var newUser models.User
	newUser.Id = bson.NewObjectId()
	newUser.FId = details.User_id
	newUser.Name = userDetails.Name
	newUser.Email = userDetails.Email
	newUser.Avatar = defaultAvatar

	return newUser
}

//Authenticates an existing user session and retrieves the
//user ID it was assigned to
func AuthenticateSession(sid string, mdb *mgo.Database) string {
	//Lookup session id

	//Check if it is valid and has not expired
	//Return uid if true
	//Otherwise return empty string
	return ""
}

func AuthHandler(c *gin.Context) {
	authToken := c.Query("auth_token")
	mdb := c.MustGet("db").(*mgo.Database)
	uid, err := AuthenticateUser(authToken, mdb)
	if err != nil {
		//Token is invalid
		panic("Token error, should probably handle this better...")
	}
	sessionKey := NewSession(uid, mdb)
	c.SetCookie(sessionCookieName, sessionKey.SessionKey, sessionDuration, "/", "apx.twintailsare.moe", false, false)
	sessionUser, err := AuthSession(c, mdb)
        if (sessionUser != nil) {
	  fmt.Printf("Session cookie found for %q\n", sessionUser)
        }
}

func AuthSession(c *gin.Context, mdb *mgo.Database) (*bson.ObjectId, error) {
	sessionKey, err := c.Cookie(sessionCookieName)
	if err != nil {
		return nil, err
		//Cookie not found, apparently...
	}

	//Lookup session in database
	query := mdb.C(models.CollectionSessions).Find(bson.M{"session_key": sessionKey})
	n, err := query.Count()
	if err != nil {
		fmt.Println("couldnt run query")
		return nil, err
	}
	if n == 0 {
		fmt.Println("no matching cookies")
		return nil, errors.New("No such session")
	}

	results := query.Iter()

	var session models.Session
	currentTime := time.Now().UTC().Unix()
	for results.Next(&session) {
		if session.Expires < currentTime {
			mdb.C(models.CollectionSessions).RemoveId(session.Id)
		} else {
                        fuser, _ := models.RetrieveUser(session.UID, mdb)
                        if (fuser != nil) {
                          fmt.Printf("Found a session belonging to %s\n",
                                      fuser.Name)
                        }
			return &session.UID, nil
		}
	}
	return nil, errors.New("No matching session")
}

func AuthenticateUser(token string, mdb *mgo.Database) (bson.ObjectId, error) {
	userDetails := GetUserDetails(token)
	if !userDetails.Is_valid {
		return "", errors.New("Token provided is not valid or has expired.")
	}

	//Lookup the user in database
	uid := userDetails.User_id
	user := models.User{}
	query := mdb.C(models.CollectionUsers).Find(bson.M{"fid": uid})
	no, err := query.Count()
	if err != nil {
		return "", err
	} else if no != 0 {
		//If the user exists, return the user id and nil
		err := query.One(&user)
		if err != nil {
			return "", err
		}
	} else {
		//Otherwise, make a new user with BuildUser, and add it to the database
		user = BuildUser(userDetails)
		mdb.C(models.CollectionUsers).Insert(user)
	}
	return user.Id, nil
}
