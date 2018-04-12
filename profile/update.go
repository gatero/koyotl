package profile

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

func Upsert(id string, update map[string]interface{}) (*mgo.ChangeInfo, error) {
	c, _ := Collection()
	var selector map[string]interface{}

	selector["_id"] = bson.ObjectIdHex(id)
	i, e := c.Upsert(selector, update)
	if e != nil {
		return nil, e
	}
	return i, nil
}

func RH_Upsert(c *gin.Context) {
	var update map[string]interface{}
	c.ShouldBindJSON(&update)
	fmt.Printf("\n\n UPSERT %s\n\n", update)

	i, e := Upsert(c.Param("id"), update)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e)
	}

	c.JSON(http.StatusOK, i)
}
