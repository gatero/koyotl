package profile

import (
	"app/db"
	"os"

	mgo "gopkg.in/mgo.v2"
)

type Profile struct {
	Id        string
	FirstName string
	Name      string
	Email     string
	LastName  string
	Birthday  string
}

var MONGO_DATABASE string
var COLLECTION string

func init() {
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	COLLECTION = "profile"
}

func ProfileC() (*mgo.Collection, *mgo.Session, error) {
	s, e := db.Mongo()
	if e != nil {
		return nil, nil, e
	}
	s.SetMode(mgo.Monotonic, true)

	return s.DB(MONGO_DATABASE).C(COLLECTION), s, nil
}
