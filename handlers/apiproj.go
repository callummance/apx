package handlers

import (
	"github.com/callummance/apx-srv/auth"
	"github.com/callummance/apx-srv/db"
	"github.com/callummance/apx-srv/models"
	"github.com/gin-gonic/gin"
)

func isProjPublic(pid string) bool {
	rdb := db.ReactSession
	_, found, err := rdb.GetProject(pid)

	if err != nil {
		return false
	} else if !found {
		return false
	} else {
		return true
	}
}

func postNewProject(c *gin.Context) {
	rdb := db.ReactSession
	uid, found, err := auth.AuthSession(c, rdb)
	if !found && err != nil {
		c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	} else if !found {
		c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	} else if err != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
		me, found, err := rdb.GetUser(uid)
		if err != nil {
			c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
		} else if !found {
			c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
		} else {
			proj, err := rdb.WriteProject(me)
			if err != nil {
				c.String(500, `{"code": -1, "message": "Could not insert a new project"}`)
			} else {
				c.JSON(201, proj)
			}
		}
	}
}

func getProject(c *gin.Context) {
	rdb := db.ReactSession
	pid := c.Param("pid")
	proj, found, err := rdb.GetProject(pid)

	if !found && err == nil {
		c.String(403, "{\"code\": 2000, \"message\": \"Could not find that project\"}")
	} else if err != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
		if proj.Public {
			c.JSON(200, proj)
		} else {
			uid, _, err := auth.AuthSession(c, rdb)
			if err != nil {
				c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
			} else if userOwnsProj(uid, proj) {
				c.JSON(201, proj)
			} else {
				c.String(403, `{"code": 2001, "message": "You need to be an owner to view that project"}`)
			}
		}
	}

}

func userOwnsProj(uid string, proj *models.Project) bool {
	for _, u := range proj.Owners {
		if u == uid {
			return true
		}
	}
	return false
}
