package profile

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

var NoIdError string = "No id present in the request"

// TODO: validate the object index
func FindById(id string) (Profile, error) {
	var p Profile
	c, _ := Collection()
	query := bson.M{"_id": bson.ObjectIdHex(id)}

	if len(id) == 0 {
		return p, errors.New(NoIdError)
	}

	if e := c.Find(query).One(&p); e != nil {
		return p, e
	}

	return p, nil
}

func RH_FindById(c *gin.Context) {
	p, e := FindById(c.Param("id"))
	if e != nil {
		if e.Error() == NoIdError {
			c.JSON(http.StatusBadRequest, e)
		}
		c.JSON(http.StatusNotFound, e)
	}

	c.JSON(http.StatusOK, p)
}
