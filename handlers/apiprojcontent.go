package handlers

import (
	"github.com/callummance/apx-srv/db"
	"github.com/callummance/apx-srv/models"
	"github.com/callummance/apx-srv/events"
	"github.com/gin-gonic/gin"
        "fmt"
)

func getProjectContent(c *gin.Context) {
	rdb := db.ReactSession
	pid := c.Param("pid")
	proj, err := rdb.GetProjectContent(pid)

	if err != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
	  c.JSON(201, proj)
	}

}

func modifyProjContent(c *gin.Context, projC *models.ProjectContentTemp) {
  rdb := db.ReactSession
  oldPID := projC.Id
  c.BindJSON(projC)
  projC.Id = oldPID
  modified, err := rdb.ModifyProjectContent(projC)
  if (err != nil) {
    c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
  } else if (!modified) {
    c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
  } else {
    c.Status(201)
  }

}

func writeProjectContent(c *gin.Context) {
	rdb := db.ReactSession
	pid := c.Param("pid")
	proj, err := rdb.GetProjectContent(pid)
	if err != nil {
  	  c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
          fmt.Println(proj)
          modifyProjContent(c, proj)
          events.UpdateProject(pid, proj.Content)
	}

}
