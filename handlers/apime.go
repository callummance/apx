package handlers


import (
  "github.com/gin-gonic/gin"
  "github.com/callummance/apx-srv/db"
  "github.com/callummance/apx-srv/auth"
)

func getMeHandler(c *gin.Context) {
  rdb := db.ReactSession

  //Get the cookie
  curUser, found, err := auth.AuthSession(c, rdb)
  if (!found && err != nil) {
    c.String(401, "{'code': 1001, 'message': 'No session key was provided'}")
  } else if (!found) {
    c.String(403, "{'code': 1000, 'message': 'Could not find that session'}")
  } else if (err != nil) {
    c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
  } else {
    me, found, err := rdb.GetUser(curUser)
    if (err != nil){
      c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
    } else if (!found) {
      c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
    } else {
      c.JSON(200, me)
    }
  }
}

func getFriendHandler(c *gin.Context) {
  rdb := db.ReactSession

  //Get the cookie
  curUser, found, err := auth.AuthSession(c, rdb)
  if (!found && err != nil) {
    c.String(401, "{'code': 1001, 'message': 'No session key was provided'}")
  } else if (!found) {
    c.String(403, "{'code': 1000, 'message': 'Could not find that session'}")
  } else if (err != nil) {
    c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
  } else {
    me, found, err := rdb.GetUser(curUser)
    if (err != nil){
      c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
    } else if (!found) {
      c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
    } else {
      c.JSON(200, me.Friends)
    }
  }
}

func getProjHandler(c *gin.Context) {
  rdb := db.ReactSession

  //Get the cookie
  curUser, found, err := auth.AuthSession(c, rdb)
  if (!found && err != nil) {
    c.String(401, "{'code': 1001, 'message': 'No session key was provided'}")
  } else if (!found) {
    c.String(403, "{'code': 1000, 'message': 'Could not find that session'}")
  } else if (err != nil) {
    c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
  } else {
    me, found, err := rdb.GetUser(curUser)
    if (err != nil){
      c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
    } else if (!found) {
      c.String(500, "{'code': -1, 'message': 'An unexpected error occurred'}")
    } else {
      c.JSON(200, me.Projects)
    }
  }
}
