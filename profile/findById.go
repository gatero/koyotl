package profile

import (
	"app/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func FindById(c *gin.Context) {
	session, _ := db.Mongo()
	defer session.Close()

	p := session.DB(MONGO_DATABASE).C(COLLECTION)
	id := bson.M{"_id": bson.ObjectIdHex(c.Param("id"))}
	profile := Profile{}
	err := p.Find(id).One(&profile)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	c.JSON(http.StatusOK, profile)
}
