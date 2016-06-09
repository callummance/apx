package handlers

import (
	"github.com/callummance/apx-srv/db"
	"github.com/callummance/apx-srv/auth"
	"github.com/callummance/apx-srv/models"
	"github.com/callummance/apx-srv/events"
	"github.com/gin-gonic/gin"
        "fmt"
)


func isSnippetPublic(sid string) bool {
	rdb := db.ReactSession
	_, found, err := rdb.GetSnippet(sid)

	if err != nil {
		return false
	} else if !found {
		return false
	} else {
		return true
	}
}

func userOwnsSnippet(uid string, snippet *models.Snippet) bool {
  return snippet.Owner == uid
}

func postNewSnippet(c *gin.Context) {
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
			snippet, err := rdb.WriteSnippet(me)
			if err != nil {
				c.String(500, `{"code": -1, "message": "Could not insert a new project"}`)
			} else {
				c.JSON(201, snippet)
			}
		}
	}
}

func getSnippet (c *gin.Context) {
	rdb := db.ReactSession
	sid := c.Param("sid")
	snippet, found, err := rdb.GetSnippet(sid)

	if !found && err == nil {
		c.String(404, "{\"code\": 2000, \"message\": \"Could not find that snippet\"}")
	} else if err != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
		if snippet.Public {
			c.JSON(200, snippet)
		} else {
			uid, _, err := auth.AuthSession(c, rdb)
			if err != nil {
				c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
			} else if userOwnsSnippet(uid, snippet) {
				c.JSON(201, snippet)
			} else {
				c.String(403, `{"code": 3001, "message": "You need to be an owner to view that snippet"}`)
			}
		}
	}

}

func postSnippetHandler(c *gin.Context, mod string) {
  rdb := db.ReactSession

  //Get the cookie
  curUser, found, err := auth.AuthSession(c, rdb)
  if (!found && err != nil) {
    c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
  } else if (!found) {
    c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
  } else if (err != nil) {
    c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
  } else {
    sid := c.Param("sid")
    snippet, found, err := rdb.GetSnippet(sid)

    if !found && err == nil {
      c.String(404, "{\"code\": 2000, \"message\": \"Could not find that snippet\"}")
    } else if err != nil {
      c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
    } else if (!userOwnsSnippet(curUser, snippet)) {
      c.String(403, "{\"code\": 2001, \"message\": \"You do not own that snippet\"}")
    } else if (mod == "/meta"){
      modifySnippetMeta(c, snippet)
    } else {
      fmt.Println(mod)
    }
  }
}


func modifySnippetMeta(c *gin.Context, snippet *models.Snippet) {
  rdb := db.ReactSession
  oldSID := snippet.Id
  oldOwner := snippet.Owner
  c.BindJSON(snippet)
  snippet.Id = oldSID
  snippet.Owner = oldOwner

  modified, err := rdb.ModifySnippet(snippet)
  if (err != nil) {
    c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
  } else if (!modified) {
    c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
  } else {
    c.Status(201)
  }
}


func getSnippetContent(c *gin.Context) {
	rdb := db.ReactSession
	sid := c.Param("sid")
	snippet, err := rdb.GetSnippetContent(sid)

	if err != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
	  c.JSON(201, snippet)
	}

}


func modifySnippetContent(c *gin.Context, snippetC *models.InstrumentSnippet) {
  rdb := db.ReactSession
  oldSID := snippetC.Id
  c.BindJSON(snippetC)
  snippetC.Id = oldSID
  modified, err := rdb.ModifySnippetContent(snippetC)
  if (err != nil) {
    c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
  } else if (!modified) {
    c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
  } else {
    c.Status(201)
  }

}

func writeSnippetContent(c *gin.Context) {
	rdb := db.ReactSession
	sid := c.Param("sid")
	snippet, err := rdb.GetProjectContent(sid)
	if err != nil {
  	  c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	} else {
          fmt.Println(snippet)
          modifyProjContent(c, snippet)
          events.UpdateSnippet(sid, snippet.Content)
	}

}
