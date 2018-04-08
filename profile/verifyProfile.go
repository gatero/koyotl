package profile

import (
	"fmt"

	"firebase.google.com/go/auth"
	bson "gopkg.in/mgo.v2/bson"
)

func VerifyProfile(t *auth.Token) {
	c, _ := Collection()

	p := Profile{
		FirebaseId: t.UID,
		Name:       t.Claims["name"].(string),
		Email:      t.Claims["email"].(string),
		Status:     STATUS_INCOMPPLETE,
		Role:       ROLE_AUTHOR,
	}

	r := Profile{}
	email := bson.M{"email": p.Email}

	if e := c.Find(email).One(&r); e != nil {
		fmt.Printf("\n\nVerify Profile ERROR: %s\n\n", e)
	}

	if (Profile{}) != r {
		return
	}

	if _, e := Create(&p); e == nil {
		fmt.Printf("New profile was created: %s", p)
	}

	//firebase := t.Claims["firebase"].(map[string]interface{})
	//identities := firebase["identities"].(map[string]interface{})
}
