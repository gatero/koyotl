package profile

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func Find(c *gin.Context) {
	session, _ := db.Mongo()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var profiles []Profile
	p := session.DB(MONGO_DATABASE).C(COLLECTION)
	p.Find(nil).All(&profiles)

	c.JSON(http.StatusOK, profiles)
}
