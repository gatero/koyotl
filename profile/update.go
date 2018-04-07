package profile

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func RH_Update(c *gin.Context) {
	session, _ := db.Mongo()
	defer session.Close()

	p := session.DB(MONGO_DATABASE).C(COLLECTION)
	id := bson.M{"_id": bson.ObjectIdHex(c.Param("id"))}

	profile := Profile{
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Birthday:  c.PostForm("birthday"),
	}

	err := p.Update(id, &profile)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, profile)
}
