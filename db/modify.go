package db

import (
	"errors"
	"fmt"
	"github.com/callummance/apx-srv/models"
	"github.com/dancannon/gorethink"
)

func (c *DbConn) GetUUID() string {
	query, err := gorethink.UUID().Run(c.Session)
	if err != nil {
		panic("Database broke again")
	} else {
		var uuid string
		err = query.One(&uuid)
		if err != nil {
			panic("Could not generate a uuid")
		} else {
			return uuid
		}
	}
}

func (c *DbConn) WriteSession(session models.Session) error {
	resp, err := SessionTable.Insert(session).RunWrite(c.Session)
	if err != nil {
		return err
	} else if resp.Errors != 0 {
		return errors.New("Database insert failed")
	} else if resp.Inserted != 1 {
		return errors.New("Incorrect number of sessions inserted")
	} else {
		return nil
	}
}

func (c *DbConn) GetProject(pid string) (*models.Project, bool, error) {
	resp, err := ProjectTable.Get(pid).Run(c.Session)
	if err != nil {
		return nil, false, err
	}
	defer resp.Close()

	//Check that a result was found
	if resp.IsNil() {
		//No results were found
		return nil, false, nil
	} else {
		session := models.Project{}
		err = resp.One(&session)
		if err != nil {
			return nil, false, err
		} else {
			return &session, true, nil
		}
	}
}


func (c *DbConn) GetProjectContent(pid string) (*models.ProjectContentTemp, error) {
	resp, err := ProjectCTable.Get(pid).Run(c.Session)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	//Check that a result was found
	if resp.IsNil() {
		//No results were found
		return nil, errors.New("wat")
	} else {
		session := models.ProjectContentTemp{}
		err = resp.One(&session)
		if err != nil {
			return nil, err
		} else {
			return &session, nil
		}
	}
}

func (c *DbConn) WriteProject(me *models.User) (*models.Project, error) {
	proj := models.NewDefaultProject(me.Id)
	proj.Id = c.GetUUID()
        projContent := models.ProjectContentTemp{}
        projContent.Id = proj.Id
        projContent.Content = ""


	me.Projects = append(me.Projects, proj.Id)
	found, err := c.ModifyUser(me)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, errors.New("Could not find that user")
	}

	resp, err := ProjectTable.Insert(proj).RunWrite(c.Session)
	_, _ = ProjectCTable.Insert(projContent).RunWrite(c.Session)
	if err != nil {
		return nil, err
	} else if resp.Errors != 0 {
		return nil, errors.New("Database insert failed")
	} else if resp.Inserted != 1 {
		return nil, errors.New("Incorrect number of projects inserted")
	} else {
		return &proj, nil
	}

}

func (c *DbConn) ModifyProjectContent(proj *models.ProjectContentTemp) (bool, error) {
  res, err := ProjectCTable.Get(proj.Id).Update(*proj).RunWrite(c.Session)
  if err != nil {
    return false, err
  } else if res.Replaced == 0 {
    return false, nil
  } else {
    return true, err
  }
}


func (c *DbConn) ModifyProject(proj *models.Project) (bool, error) {
  res, err := ProjectTable.Get(proj.Id).Update(*proj).RunWrite(c.Session)
  if err != nil {
    return false, err
  } else if res.Replaced == 0 {
    return false, nil
  } else {
    return true, err
  }
}

func (c *DbConn) GetSession(sessionKey string) (*models.Session, bool, error) {
	query, err := SessionTable.Get(sessionKey).Run(c.Session)
	if err != nil {
		fmt.Printf("Could not find session for key %s\n", sessionKey)
		return nil, false, err
	}

	defer query.Close()

	//Check that a result was found
	if query.IsNil() {
		//No results were found
		fmt.Printf("Could not find session for key %s\n", sessionKey)
		return nil, false, nil
	} else {
		foundSession := models.Session{}
		err = query.One(&foundSession)
		if err != nil {
			return nil, false, err
		} else {
			return &foundSession, true, nil
		}
	}
}

func (c *DbConn) ModifyUser(user *models.User) (bool, error) {
	res, err := UserTable.Get(user.Id).Update(*user).RunWrite(c.Session)
	if err != nil {
		return false, err
	} else if res.Replaced == 0 {
		return false, nil
	} else {
		return true, err
	}
}

func (c *DbConn) GetUser(uid string) (*models.User, bool, error) {
	query, err := UserTable.Get(uid).Run(c.Session)
	if err != nil {
		return nil, false, err
	}

	defer query.Close()

	//Check that a result was found
	if query.IsNil() {
		//No results were found
		return nil, false, nil
	} else {
		foundUser := models.User{}
		err = query.One(&foundUser)
		if err != nil {
			return nil, false, err
		} else {
			return &foundUser, true, nil
		}
	}
}

func (c *DbConn) WriteUser(user models.User) error {
	resp, err := UserTable.Insert(user).RunWrite(c.Session)
	if err != nil {
		return err
	} else if resp.Errors != 0 {
		return errors.New("Database insert failed")
	} else if resp.Inserted != 1 {
		return errors.New("Incorrect number of users inserted")
	} else {
		return nil
	}
}

func (c *DbConn) GetFBUser(fid string) (*models.User, bool, error) {
	fmt.Printf("Authenticating FID %s\n", fid)
	query, err := UserTable.Filter(map[string]interface{}{"fid": fid}).Run(c.Session)
	if err != nil {
		fmt.Println("Failed to search db")
		return nil, false, err
	}

	defer query.Close()

	//Check that a result was found
	if query.IsNil() {
		fmt.Println("Failed to find user")
		//No results were found
		return nil, false, nil
	} else {
		foundUser := models.User{}
		err = query.One(&foundUser)
		if err != nil {
			fmt.Printf("There was an error: %s\n", err)
			return nil, false, err
		} else {
			return &foundUser, true, nil
		}
	}
}

func (c *DbConn) RemoveUserFromProject(pid string, uid string) (bool, error) {
  res, err := UserTable.Get(uid).Update(map[string]interface{}{
    "projects": gorethink.Row.Field("projects").SetDifference([]string{pid}),
  }).RunWrite(c.Session)
  if err != nil {
    return false, err
  } else if res.Replaced == 0 {
    return false, nil
  } else {
    res, err := ProjectTable.Get(pid).Update(map[string]interface{}{
      "owner": gorethink.Row.Field("owner").SetDifference([]string{uid}),
    }).RunWrite(c.Session)
    if err != nil {
      return false, err
    } else if res.Replaced == 0 {
      return false, nil
    } else {
      res, err := ProjectTable.Get(pid).Field("owner").Count().Run(c.Session)
      defer res.Close()
      if err != nil {
        return false, err
      } else {
        var cnt int
        res.One(&cnt)
        if cnt == 0 {
          _, _ = ProjectTable.Get(pid).Delete().RunWrite(c.Session)
        }
        return true, nil
      }
    }
  }
}
