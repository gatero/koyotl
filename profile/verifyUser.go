package profile

import (
	"app/db"
	"fmt"

	"firebase.google.com/go/auth"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

func VerifyUser(decodedToken *auth.Token) {
	session, _ := db.Mongo()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	p := session.DB(MONGO_DATABASE).C(COLLECTION)

	profile := &Profile{
		Id:         bson.NewObjectId(),
		FirebaseId: bson.ObjectIdHex(decodedToken.UID),
		Name:       decodedToken.Claims["firstName"].(string),
		FirstName:  "",
		LastName:   "",
		Email:      decodedToken.Claims["email"].(string),
		Birthday:   "",
	}

	email := bson.M{"Email": profile.Email}
	result := Profile{}
	err := p.Find(email).One(&result)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	if (Profile{}) == result {
		return
	}

	err = p.Insert(profile)
	if err != nil {
		panic(err)
	}

	//firebase := decodedToken.Claims["firebase"].(map[string]interface{})
	//identities := firebase["identities"].(map[string]interface{})
}
