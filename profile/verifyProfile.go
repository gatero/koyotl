package profile

import (
	"fmt"

	"firebase.google.com/go/auth"
	bson "gopkg.in/mgo.v2/bson"
)

func VerifyProfile(t *auth.Token) {
	c, e := ProfileC()
	if e != nil {
		panic(e)
	}

	fmt.Printf("\n\n VerifyProfile TOKEN %s\n\n", t)

	p := Profile{
		Id:    t.UID,
		Name:  t.Claims["name"].(string),
		Email: t.Claims["email"].(string),
	}

	r := Profile{}
	email := bson.M{"email": p.Email}
	err := c.Find(email).One(&r)

	if err != nil {
		fmt.Printf("\n\nVerify Profile ERROR: %s\n\n", err)
	}

	if (Profile{}) != r {
		return
	}

	profile, err := Create(&p)
	if err == nil {
		fmt.Printf("New profile was created: %s", profile)
	}

	//firebase := t.Claims["firebase"].(map[string]interface{})
	//identities := firebase["identities"].(map[string]interface{})
}
