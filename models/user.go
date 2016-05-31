package models

import (
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUsers = "users"
        CollectionSessions = "session"
)

type User struct {
  Id       string   `gorethink:"id,omitempty" json:"-"`
  FId      string   `gorethink:"fid" json:"-"`
  Name     string   `gorethink:"name" json:"name"`
  Avatar   string   `gorethink:"avatar" json:"avatar"`
  Email    string   `gorethink:"email" json:"email"`
  Projects []string `gorethink:"projects" json:"-"`
  Friends  []string `gorethink:"friends" json:"-"`
}

type Session struct {
  SessionKey string   `gorethink:"id"`
  UID        string   `gorethink:"uid"`
  Expires    int64    `gorethink:"expires_at"`
}

func (user User)NewDefaultProject() Project {
	project := Project{
		Public:true,
		Name:"untitled",
		Owners:[]string{user.Id},
	}
	return project
}

