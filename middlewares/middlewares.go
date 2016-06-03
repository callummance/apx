// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
  "net/http"

  "github.com/callummance/apx-srv/db"
  "github.com/callummance/apx-srv/auth"
  "github.com/gin-gonic/gin"
  "fmt"
)

func AuthMiddleware(c *gin.Context) {
  if (c.Request.URL.Path != "/dashboard") {
    c.Next()
    return
  }
  rdb := db.ReactSession
  _, found, err := auth.AuthSession(c, rdb)
  if !found {
    c.Redirect(307, "/login")
    fmt.Println("User not logged in")
    c.Abort()
  } else if err != nil {
    panic("wat")
  } else {
    c.Next()
  }
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
  c.Next()

  // TODO: Handle it in a better way
  if len(c.Errors) > 0 {
          c.HTML(http.StatusBadRequest, "400", gin.H{
                  "errors": c.Errors,
          })
  }
}

func CORS(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "*")
  c.Next()
}
