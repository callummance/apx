package events

import (
  "github.com/gin-gonic/gin"
  "io"
)


func GetSubP(c *gin.Context) {
  pid := c.Param("pid")
  listener := SubscribeToProject(pid)
  defer UnsubscribeFromProject(pid, listener)

  c.Stream(func(w io.Writer) bool {
    c.SSEvent("update", <-listener)
    return true
  })
}

func GetSubS(c *gin.Context) {
  sid := c.Param("sid")
  listener := SubscribeToSnippet(sid)
  defer UnsubscribeFromProject(sid, listener)

  c.Stream(func(w io.Writer) bool {
    c.SSEvent("update", <-listener)
    return true
  })
}
