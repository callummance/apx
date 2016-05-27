package models

import (
  "gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2"
  "errors"
  "fmt"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUsers = "users"
        CollectionSessions = "session"
)

type User struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
        FId      string `json:"fid" binding:"required" bson:"fid"`
	Name     string `json:"name" binding:"required" bson:"name"`
	Avatar   string `json:"avatar" binding:"required" bson:"avatar"`
        Email    string `json: "email" binding:"required" bson:"email"`
	Projects []string `json:"projects" bson:"projects"`
}

type Session struct {
        Id         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
        UID        bson.ObjectId `json:"uid" binding:"required" bson:"uid"`
        SessionKey string `json:"session_key" binding:"required" bson:"session_key"`
        Expires    int64  `json:"expires_at" bingin:"required" bson:"expires_at"`
}

func (user User)NewDefaultProject() Project {
	project := Project{
		Public:true,
		Name:"untitled",
		Owners:[]bson.ObjectId{user.Id},
	}
	return project
}

func RetrieveUser(uid bson.ObjectId, mdb *mgo.Database) (*User, error) {
  query := mdb.C(CollectionUsers).FindId(uid)
  n, err := query.Count()
  if (err != nil) {
    panic("Aaaaargh database broke again!")
  } else if (n > 1) {
    panic("Whoops a user was in the database more than once lul")
  } else if (n == 0) {
    fmt.Println("Found no users")
    return nil, errors.New("Could not find a user with that ID")
  } else {
    foundUser := User{}
    err = query.One(&foundUser)
    if (err != nil) {
      return nil, err
    } else {
      return &foundUser, nil 
    }
  }
}
