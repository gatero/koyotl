package profile

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

func Create(c *gin.Context) {
	session, _ := db.Mongo()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	p := session.DB(MONGO_DATABASE).C(COLLECTION)

	profile := &Profile{
		Id:        bson.NewObjectId(),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Birthday:  c.PostForm("birthday"),
	}

	err := p.Insert(profile)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, profile)
}
