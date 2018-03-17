package profile

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func DeleteById(c *gin.Context) {
	session, _ := db.Mongo()
	defer session.Close()

	p := session.DB(MONGO_DATABASE).C(COLLECTION)
	id := bson.M{"_id": bson.ObjectIdHex(c.Param("id"))}

	err := p.Remove(id)

	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, "The profile %s was deleted successfully !", c.Param("id"))
}
