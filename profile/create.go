package profile

import (
	"app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func Create(p *Profile) (*Profile, error) {
	s, e := db.Mongo()
	if e != nil {
		panic(e)
	}
	defer s.Close()
	s.SetMode(mgo.Monotonic, true)

	c := s.DB(MONGO_DATABASE).C(COLLECTION)

	if e := c.Insert(p); e != nil {
		return nil, e
	}

	return p, nil
}

func RH_Create(c *gin.Context) {
	var p Profile
	if e := c.ShouldBindJSON(&p); e == nil {
		if p, e := Create(&p); e == nil {
			c.JSON(http.StatusOK, p)
		}
	}
}
