package handlers

import (
	"github.com/gin-gonic/gin"
        "github.com/callummance/apx-srv/events"
)

func ApiHandlers(router *gin.Engine) {
	//Handlers for '/me/*' endpoints
	router.GET("/api/me", getMeHandler)
	router.GET("/api/me/following", getFriendHandler)
	router.GET("/api/me/projects", getProjHandler)
	router.GET("/api/me/projects/meta", getProjMetaHandler)
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
	router.GET("/api/user/:uid/*action", func(c *gin.Context) {
		if c.Param("action") == "projects" {
			getPublicProj(c)
		}
	})

	//Handlers for '/project/*' endpoints
	router.POST("/api/project", postNewProject)
	router.GET("/api/project/:pid", getProject)
        router.POST("/api/project/:pid/*action", func(c *gin.Context) {
          postProjHandler(c, c.Param("action"))
        })
        router.GET("/api/projectcontent/:pid", getProjectContent)
        router.POST("/api/projectcontent/:pid", writeProjectContent)
        router.GET("/api/projectsub/:pid", events.GetSub)
}
