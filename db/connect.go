package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
        "github.com/dancannon/gorethink"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo

        //React session details
        ReactSession *DbConn

        SessionTable gorethink.Term = gorethink.DB("apx").Table("session")
        UserTable gorethink.Term = gorethink.DB("apx").Table("user")
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://db.twintailsare.moe:27017"
        
        //RethinkDB Parameters
        RethinkDBUrl = "db.twintailsare.moe:28015"
        RethinkDBDatabse = "apx"
        MaxIdle = 10
        MaxOpen = 10
)

type DbConn struct {
  Session *gorethink.Session
  ConnectionURL string
}


func ReactConnect (connectURL string, database string) *DbConn {
  var newConn DbConn
  var err error
  newConn.ConnectionURL = connectURL
  newConn.Session, err = gorethink.Connect(gorethink.ConnectOpts {
    Address: connectURL,
    Database: database,
    MaxIdle: MaxIdle,
    MaxOpen: MaxOpen,
  })
  if (err != nil) {
    panic("Could not connect to React Database. Now exiting.\n")
  }
  return &newConn
}
  

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
