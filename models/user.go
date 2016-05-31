package models

import (
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUsers = "users"
        CollectionSessions = "session"
)

type User struct {
  Id       string   `gorethink:"id,omitempty" bson:"_id,omitempty"`
  FId      string   `gorethink:"fid" bson:"fid"`
  Name     string   `gorethink:"name" bson:"name"`
  Avatar   string   `gorethink:"avatar" bson:"avatar"`
  Email    string   `gorethink:"email" bson:"email"`
  Projects []string `gorethink:"projects" projects" bson:"projects"`
}

type Session struct {
  SessionKey string   `gorethink:"id" bson:"session_key"`
  UID        string   `gorethink:"uid" bson:"uid"`
  Expires    int64    `gorethink:"expires_at" bson:"expires_at"`
}

func (user User)NewDefaultProject() Project {
	project := Project{
		Public:true,
		Name:"untitled",
		Owners:[]string{user.Id},
	}
	return project
}

