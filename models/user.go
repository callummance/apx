package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUsers = "users"
)

type User struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name" binding:"required" bson:"name"`
	Avatar   string `json:"avatar" binding:"required" bson:"avatar"`
	Projects []string `json:"projects" bson:"projects"`
}

func (user User)NewDefaultProject() Project {
	project := Project{
		Public:true,
		Name:"untitled",
		Owners:[]bson.ObjectId{user.Id},
	}
	return project
}