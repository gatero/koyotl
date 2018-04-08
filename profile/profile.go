package profile

import (
	"app/db"
	"os"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id         bson.ObjectId `form:"id" json:"id,omitempty" bson:"_id,omitempty"`
	FirebaseId string        `form:"firebaseid" json:"firebaseid,omitempty" bson:"firebaseid"`
	FirstName  string        `form:"firstname" json:"firstname,omitempty" bson:"firstname"`
	LastName   string        `form:"lastname" json:"lastname,omitempty" bson:"lastname"`
	Name       string        `form:"name,omitempty" json:"name,omitempty" bson:"name" `
	Email      string        `form:"email,omitempty" json:"email,omitempty" bson:"email`
	Birthday   string        `form:"birthday" json:"birthday,omitempty" bson:"birthday"`
	Status     string        `form:"status" json:"status,omitempty" bson:"status"`
	Role       string        `form:"role" json:"role,omitempty" bson:"role"`
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
