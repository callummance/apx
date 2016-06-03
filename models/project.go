package models

type Project struct {
  Id         string   `gorethink:"id" json:"pid"`
  Public     bool     `gorethink:"public" json:"public"`
  Name       string   `gorethink:"name" json:"name"`
  Owners     []string `gorethink:"owner" json:"owners"`
}

type ProjectContentTemp struct {
  Id string `gorethink:"id" json "pid"`
  Content string `gorethink:"content" json:"content"`
}

type ProjectContent struct {
  Id     string  `gorethink:"id" bson:"_id,omitempty"`
  Tracks []Track `gorethink:"tracks" bson:"tracks"`
  Tempo  float64 `gorethink:"tempo" bson:"tempo"`
}
type Track struct {
  Id        string  `gorethink:"id" bson:"_id,omitempty"`
  Volume    int8    `gorethink:"volume"`
  TrackType string  `gorethink:"trackType"`
  Solo      bool    `gorethink:"solo"`
  Mute      bool    `gorethink:"mute"`
}

type InstrumentTrack struct {

}

type Clip struct {
	Id        string  `json:"id"`
	StartTime int     `json:"startTime"`
	Duration  int     `json:"duration"`
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

func NewDefaultProject(uid string) Project {
  project := Project{
    Public:true,
    Name:"untitled",
    Owners:[]string{uid},
  }
  return project
}
