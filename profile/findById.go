package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func FindById(id string, p *Profile) error {
	c, _ := Collection()
	query := bson.M{"_id": bson.ObjectIdHex(id)}

	if e := c.Find(query).One(&p); e != nil {
		return e
	}

	return nil
}

func RH_FindById(c *gin.Context) {
	p := Profile{}
	id := c.Param("id")

	if e := FindById(id, &p); e != nil {
		c.JSON(http.StatusInternalServerError, e)
	}

	c.JSON(http.StatusOK, p)
}
