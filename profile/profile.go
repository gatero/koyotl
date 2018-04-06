package profile

import (
	"os"

	bson "gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	FirebaseId bson.ObjectId `bson:"firebaseId,omitempty"`
	FirstName  string
	Name       string
	Email      string
	LastName   string
	Birthday   string
}

var MONGO_DATABASE string
var COLLECTION string

func init() {
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	COLLECTION = "profile"
}
