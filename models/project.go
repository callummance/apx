package models

import "gopkg.in/mgo.v2/bson"

type Project struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Public    bool `json:"public"  binding:"required" bson:"public"`
	Name      string `json:"name" binding:"required" bson:"name"`
	Owners    []bson.ObjectId `json:"owners" binding:"required" bson:"owners"`
	ContentId  bson.ObjectId `json:"contentId" binding:"required" bson:"contentId"`
}

type ProjectContent struct {
	Id     bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Tracks []Track `json:"tracks" bson:"tracks"`
	Tempo  float64 `json:"tempo" bson:"tempo"`
}
type Track struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Volume    int8 `json:"volume"`
	TrackType string `json:"tracktype"`
	Solo      bool `json:"solo"`
	Mute      bool `json:"mute"`
}

type InstrumentTrack struct {

}

type Clip struct {
	Id        bson.ObjectId `json:"id"`
	StartTime int `json:"startTime"`
	Duration  int `json:"duration"`
}

type InstrumentClip struct {

}

type Note struct {
	Pitch     string `json:"pitch"`
	StartTime int `json:"startTime"`
	Duration  int `json:"duration"`
}

type Error struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func NewDefaultProjectContent() ProjectContent {
	projectContent := ProjectContent{
		Tracks:[]Track{},
		Tempo:120,
	}
	return projectContent
}





