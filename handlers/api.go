package handlers

import (
  "github.com/gin-gonic/gin"
)

func ApiHandlers(router *gin.Engine) {
  router.GET("/api/me", getMeHandler)
  router.GET("/api/me/friends", getFriendHandler)
  router.GET("/api/me/projects", getProjHandler)
}


