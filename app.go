package main

import (
	"bitbucket.org/callummance/apx/auth"
	"github.com/gin-gonic/gin"
	"https://gitlab.doc.ic.ac.uk/cm3914/apx.git"
	"https://gitlab.doc.ic.ac.uk/cm3914/apx.git/app/db"
	"https://gitlab.doc.ic.ac.uk/cm3914/apx.git/app/handlers/users"
	"https://gitlab.doc.ic.ac.uk/cm3914/apx.git/app/middlewares"
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
