package profile

import (
	"app/mongo"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var MONGO_DATABASE string
var COLLECTION string

func init() {
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	COLLECTION = "profile"
}

func Collection() (*mgo.Collection, error) {
	s, e := mongo.Session()
	if e != nil {
		return nil, e
	}
	s.SetMode(mgo.Monotonic, true)

	return s.DB(MONGO_DATABASE).C(COLLECTION), nil
}

type RPC struct{}
