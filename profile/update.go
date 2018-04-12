package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func Upsert(id string, update map[string]interface{}) error {
	c, _ := Collection()
	var selector map[string]interface{}

	selector["_id"] = bson.ObjectIdHex(id)
	if e := c.Upsert(selector, update); e != nil {
		return e
	}
	return nil
}

func RH_Upsert(c *gin.Context) {
	var update map[string]interface{}
	c.ShouldBindJSON(&update)

	if e := Upsert(c.Param("id"), update); e != nil {
		c.JSON(http.StatusInternalServerError, e)
	}

	c.JSON(http.StatusOK, update)
}
