package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/callummance/apx-srv/db"
  //"github.com/callummance/apx-srv/auth"
  //"github.com/callummance/apx-srv/models"
)

func getUser (c *gin.Context) {
  rdb := db.ReactSession

  targetId := c.Param("uid")
  target, found, err := rdb.GetUser(targetId)
  if (err != nil){
    c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
  } else if (!found) {
    c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
  } else if (target.Private) {
    c.String(403, "{\"code\": 1004, \"message\": \"User has set their profile to private\"}")
  } else {
    c.JSON(200, target)
    c.JSON(200, target)
    c.JSON(200, target)
    c.JSON(200, target)
    c.JSON(200, target)
  }
}

func getPublicProj(c *gin.Context) {
  rdb := db.ReactSession

  targetId := c.Param("uid")
  target, found, err := rdb.GetUser(targetId)
  if (err != nil){
    c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
  } else if (!found) {
    c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
  } else if (target.Private) {
    c.String(403, "{\"code\": 1004, \"message\": \"User has set their profile to private\"}")
  } else {
    c.JSON(200, target.Projects)
  }
}
