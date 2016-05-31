package main

import (
	"github.com/callummance/apx-srv/auth"
	"github.com/callummance/apx-srv/db"
	"github.com/callummance/apx-srv/handlers"
	"github.com/callummance/apx-srv/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MongoDB : apx.twintailsare.moe

func main() {

        //Connect to the database and start gin router
	db.Connect()
        db.ReactSession = db.ReactConnect(db.RethinkDBUrl, db.RethinkDBDatabse)
	router := gin.Default()

	// Middlewares
	//router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

        //Serve static angular files
	router.StaticFS("/home", http.Dir("../../../../../assignments/apex_name_subject_to_change/webpage"))

        //Redirect for logged-in users
	router.GET("/loggedin", func(c *gin.Context) {
		c.Redirect(303, "/home/dashboard.html")
	})

        //Login page which installs the session cookie
	router.GET("/fbauth", auth.AuthHandler)

        //Redirect for landing page
	router.GET("/", handlers.LandingHandler)

        //Run the server
	router.Run(":80")

}
