package profile

import (
	pb "app/grpc"
	"fmt"

	"firebase.google.com/go/auth"
	bson "gopkg.in/mgo.v2/bson"
)

func VerifyProfile(t *auth.Token) {
	c, _ := Collection()

	p := pb.Profile{
		FirebaseId: t.UID,
		Name:       t.Claims["name"].(string),
		Email:      t.Claims["email"].(string),
		Status:     STATUS_INCOMPPLETE,
		Role:       ROLE_AUTHOR,
	}

	r := pb.Profile{}
	email := bson.M{"email": p.Email}

	if e := c.Find(email).One(&r); e != nil {
		fmt.Printf("\n\nVerify Profile ERROR: %s\n\n", e)
	}

	if (pb.Profile{}) != r {
		return
	}

	if e := Create(&p); e == nil {
		fmt.Printf("New profile was created: %s", p)
	}

	//firebase := t.Claims["firebase"].(map[string]interface{})
	//identities := firebase["identities"].(map[string]interface{})
}
