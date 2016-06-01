package handlers

import (
  "github.com/gin-gonic/gin"
)

func ApiHandlers(router *gin.Engine) {
  //Handlers for '/me/*' endpoints
  router.GET("/api/me", getMeHandler)
  router.GET("/api/me/friends", getFriendHandler)
  router.GET("/api/me/projects", getProjHandler)

  router.POST("/api/me", func(c *gin.Context) {
    postMeHandler(c, "general")
  })
  router.POST("/api/me/addfriend", func(c *gin.Context) {
    postMeHandler(c, "addfriend")
  })
  router.POST("/api/me/removefriend", func(c *gin.Context) {
    postMeHandler(c, "removefriend")
  })
}


