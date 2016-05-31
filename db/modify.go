package db

import (
  "github.com/dancannon/gorethink"
  "github.com/callummance/apx-srv/models"
  "errors"
  "fmt"
)


func (c *DbConn) GetUUID() string {
  query, err := gorethink.UUID().Run(c.Session)
  if (err != nil) {
    panic("Database broke again")
  } else {
    var uuid string
    err = query.One(&uuid)
    if (err != nil) {
      panic("Could not generate a uuid")
    } else {
      return uuid
    }
  }
}

func (c *DbConn) WriteSession(session models.Session) error {
  resp, err := SessionTable.Insert(session).RunWrite(c.Session)
  if (err != nil) {
    return err
  } else if (resp.Errors != 0) {
    return errors.New("Database insert failed")
  } else if (resp.Inserted != 1) {
    return errors.New("Incorrect number of sessions inserted")
  } else {
    return nil
  }
}

func (c *DbConn) GetSession(sessionKey string) (*models.Session, bool, error) {
  query, err := SessionTable.Get(sessionKey).Run(c.Session)
  if (err != nil) {
    fmt.Printf("Could not find session for key %s\n", sessionKey)
    return nil, false, err
  }

  defer query.Close()

  //Check that a result was found
  if (query.IsNil()) {
    //No results were found
    fmt.Printf("Could not find session for key %s\n", sessionKey)
    return nil, false, nil
  } else {
    foundSession := models.Session{}
    err = query.One(&foundSession)
    if (err != nil) {
      return nil, false, err
    } else {
      return &foundSession, true, nil
    }
  }
}

func (c *DbConn) GetUser(uid string) (*models.User, bool, error) {
  query, err := UserTable.Get(uid).Run(c.Session)
  if (err != nil) {
    return nil, false, err
  }

  defer query.Close()

  //Check that a result was found
  if (query.IsNil()) {
    //No results were found
    return nil, false, nil
  } else {
    foundUser := models.User{}
    err = query.One(&foundUser)
    if (err != nil) {
      return nil, false, err
    } else {
      return &foundUser, true, nil
    }
  }
}

func (c *DbConn) WriteUser(user models.User) error {
  resp, err := UserTable.Insert(user).RunWrite(c.Session)
  if (err != nil) {
    return err
  } else if (resp.Errors != 0) {
    return errors.New("Database insert failed")
  } else if (resp.Inserted != 1) {
    return errors.New("Incorrect number of users inserted")
  } else {
    return nil
  }
}

func (c *DbConn) GetFBUser(fid string) (*models.User, bool, error) {
  fmt.Printf("Authenticating FID %s\n", fid)
  query, err := UserTable.Filter(map[string]interface{}{"fid": fid}).Run(c.Session)
  if (err != nil) {
    fmt.Println("Failed to search db")
    return nil, false, err
  }

  defer query.Close()

  //Check that a result was found
  if (query.IsNil()) {
    fmt.Println("Failed to find user")
    //No results were found
    return nil, false, nil
  } else {
    foundUser := models.User{}
    err = query.One(&foundUser)
    if (err != nil) {
      fmt.Printf("There was an error: %s\n", err)
      return nil, false, err
    } else {
      return &foundUser, true, nil
    }
  }
}
