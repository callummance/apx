package db

import (
	"github.com/dancannon/gorethink"
)

var (
	//React session details
	ReactSession *DbConn

	SessionTable gorethink.Term = gorethink.DB("apx").Table("session")
	UserTable    gorethink.Term = gorethink.DB("apx").Table("user")
	ProjectTable gorethink.Term = gorethink.DB("apx").Table("project")
	ProjectCTable gorethink.Term = gorethink.DB("apx").Table("projectcontent")
)

const (
	//RethinkDB Parameters
	RethinkDBUrl     = "db.twintailsare.moe:28015"
	RethinkDBDatabse = "apx"
	MaxIdle          = 10
	MaxOpen          = 10
)

type DbConn struct {
	Session       *gorethink.Session
	ConnectionURL string
}

func ReactConnect(connectURL string, database string) *DbConn {
	var newConn DbConn
	var err error
	newConn.ConnectionURL = connectURL
	newConn.Session, err = gorethink.Connect(gorethink.ConnectOpts{
		Address:  connectURL,
		Database: database,
		MaxIdle:  MaxIdle,
		MaxOpen:  MaxOpen,
	})
	if err != nil {
		panic("Could not connect to React Database. Now exiting.\n")
	}
	return &newConn
}

// Connect connects to rethinkDB
func Connect() {
	ReactSession = ReactConnect(RethinkDBUrl, RethinkDBDatabse)
}
