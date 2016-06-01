package handlers

import (
  "github.com/gin-gonic/gin"
)

func ApiHandlers(router *gin.Engine) {
  //Handlers for '/me/*' endpoints
  router.GET("/api/me", getMeHandler)
  router.GET("/api/me/following", getFriendHandler)
  router.GET("/api/me/projects", getProjHandler)
  router.POST("/api/me", func(c *gin.Context) {
    postMeHandler(c, "general")
  })
  router.POST("/api/me/follow", func(c *gin.Context) {
    postMeHandler(c, "addfriend")
  })
  router.POST("/api/me/unfollow", func(c *gin.Context) {
    postMeHandler(c, "removefriend")
  })

  //Handlers for '/user/*' endpoints
  router.GET("/api/user/:uid", getUser)
}


