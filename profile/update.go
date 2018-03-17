package profile

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func Update(c *gin.Context) {
	session, _ := db.Mongo()
	defer session.Close()

	p := session.DB(MONGO_DATABASE).C(COLLECTION)

	profile := Profile{
		Id:        bson.ObjectIdHex(c.PostForm("id")),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Birthday:  c.PostForm("birthday"),
	}

	id := bson.M{"_id": profile.Id}

	err := p.Update(id, &profile)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, profile)
}
