package main

import (
	"github.com/callummance/apx-srv/db"
	"github.com/callummance/apx-srv/handlers/users"
	"github.com/callummance/apx-srv/handlers"
	"github.com/callummance/apx-srv/middlewares"
	"github.com/callummance/apx-srv/auth"
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


  //router.GET("/", func(c *gin.Context) {
  //  http.ServeFile(c.Writer, c.Request, "index.html")
  //})
  router.StaticFS("/home", http.Dir("../../../../../assignments/apex_name_subject_to_change/webpage"))
  router.StaticFS("/loggedin", http.Dir("./"))

  router.GET("/index.html", handlers.LandingHandler)

  router.GET("/users/:_id", users.Get)
  router.GET("/fbauth", auth.AuthHandler)

  router.StaticFS("/app/", http.Dir("app"))
  router.StaticFS("/node_modules/", http.Dir("node_modules"))

  router.GET("/", handlers.LandingHandler)

  router.Run(":80")

}
