package profile

import (
	"app/db"
	"os"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	FirebaseId string
	FirstName  string
	LastName   string
	Name       string
	Email      string
	Birthday   string
	Status     string
	Role       string
}

var MONGO_DATABASE string
var COLLECTION string

func init() {
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	COLLECTION = "profile"
}

func Collection() (*mgo.Collection, error) {
	s, e := db.Mongo()
	if e != nil {
		return nil, e
	}
	s.SetMode(mgo.Monotonic, true)

	return s.DB(MONGO_DATABASE).C(COLLECTION), nil
}

type Action struct{}

type RPC struct{}
