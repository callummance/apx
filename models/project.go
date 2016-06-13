package models

type Project struct {
	Id     string   `gorethink:"id" json:"pid"`
	Public bool     `gorethink:"public" json:"public"`
	Name   string   `gorethink:"name" json:"name"`
	Owners []string `gorethink:"owner" json:"owners"`
}

type ProjectContentTemp struct {
	Id      string `gorethink:"id" json "pid"`
	Content string `gorethink:"content" json:"content"`
}

type ProjectContent struct {
	Id     string   `gorethink:"id" bson:"_id,omitempty"`
	Tracks []string `gorethink:"tracks" bson:"tracks"`
	Tempo  int      `gorethink:"tempo" bson:"tempo"`
}

type Track struct {
	Id        string `gorethink:"id" bson:"_id,omitempty"`
	Volume    int8   `gorethink:"volume"`
	TrackType string `gorethink:"trackType"`
	Solo      bool   `gorethink:"solo"`
	Mute      bool   `gorethink:"mute"`
}

type Snippet struct {
	Id     string `gorethink:"id" json:"id"`
	Public bool   `gorethink:"public" json:"public"`
	Name   string `gorethink:"name" json:"name"`
	Owner  string `gorethink:"owner" json:"owners"`
}

type SnippetContent struct {
	Id        string `gorethink:"id" json:"id"`
	SoundFile string `gorethink:"soundfile" json:"soundfile"`
	Notes     [][]Note `gorethink:"notes" json:"notes"`
}

type Note struct {
	Pitch     string `json:"pitch"`
	Duration  string `json:"duration"`
	StartTime int    `json:"starttime"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewDefaultProjectContent() ProjectContent {
	projectContent := ProjectContent{
		Tracks: []string{},
		Tempo:  60,
	}
	return projectContent
}

func NewDefaultProject(uid string) Project {
	project := Project{
		Public: true,
		Name:   "untitled",
		Owners: []string{uid},
	}
	return project
}

func NewDefaultSnippet(uid string) Snippet {
	snippet := Snippet{
		Public: true,
		Name:   "untitled",
		Owner:  uid,
	}
	return snippet
}
