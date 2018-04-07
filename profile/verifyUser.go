package profile

import (
	"app/db"
	"fmt"

	"firebase.google.com/go/auth"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

func VerifyToken(t *auth.Token) {
	fmt.Printf("\n\nTOKEN: %s\n\n", t)
	s, e := db.Mongo()
	if e != nil {
		panic(e)
	}
	s.SetMode(mgo.Monotonic, true)

	c := s.DB(MONGO_DATABASE).C(COLLECTION)

	p := Profile{
		Id:    t.UID,
		Name:  t.Claims["name"].(string),
		Email: t.Claims["email"].(string),
	}

	fmt.Printf("decoded: %s", t)

	r := Profile{}
	email := bson.M{"Email": p.Email}
	err := c.Find(email).One(&r)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	if (Profile{}) == r {
		return
	}

	if p, err := Create(&p); err == nil {
		fmt.Printf("New profile was created: %s", p)
	}

	//firebase := t.Claims["firebase"].(map[string]interface{})
	//identities := firebase["identities"].(map[string]interface{})
}
