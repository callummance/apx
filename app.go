package main

import (
	"bitbucket.org/callummance/apx/auth"
	"github.com/callummance/apx-srv/db"
	"github.com/callummance/apx-srv/handlers/users"
	"github.com/callummance/apx-srv/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MongoDB : apx.twintailsare.moe

func main() {

	db.Connect()
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	router.GET("/users/:_id", users.Get)

	router.StaticFS("/app/", http.Dir("app"))
	router.StaticFS("/node_modules/", http.Dir("node_modules"))

	router.Run()

}
