package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/callummance/apx-srv/auth"
  "github.com/callummance/apx-srv/db"
)

func LandingHandler(c *gin.Context) {
  rdb := db.ReactSession
  uid, err := auth.AuthSession(c, rdb)
  if uid == "" {
    c.Redirect(307, "/home")
  } else if err != nil {
    panic("wat")
  } else {
    c.Redirect(307, "loggedin")
  }
}
