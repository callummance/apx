package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/callummance/apx-srv/auth"
  "gopkg.in/mgo.v2"
)

func LandingHandler(c *gin.Context) {
  mdb := c.MustGet("db").(*mgo.Database)
  uid, err := auth.AuthSession(c, mdb)
  if uid == nil {
    c.Redirect(307, "/home")
  } else if err != nil {
    panic("wat")
  } else {
    c.Redirect(307, "loggedin")
  }
}
