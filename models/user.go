package models

import "gopkg.in/mgo.v2/bson"

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
