package events

import (
  "github.com/gin-gonic/gin"
  "io"
)


func GetSub(c *gin.Context) {
  pid := c.Param("pid")
  listener := SubscribeToProject(pid)
  defer UnsubscribeFromProject(pid, listener)

  c.Stream(func(w io.Writer) bool {
    c.SSEvent("update", <-listener)
    return true
  })
}
