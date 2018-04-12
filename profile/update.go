package profile

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

func Upsert(id string, upsert interface{}) (*mgo.ChangeInfo, error) {
	c, _ := Collection()

	selector := bson.M{
		"_id": bson.ObjectIdHex(id),
	}

	i, e := c.Upsert(selector, upsert)
	if e != nil {
		return nil, e
	}
	return i, nil
}

func RH_Upsert(c *gin.Context) {
	var upsert map[string]interface{}
	c.ShouldBindJSON(&upsert)
	fmt.Printf("\n\n UPSERT %s\n\n", upsert)

	i, e := Upsert(c.Param("id"), upsert)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e)
	}

	c.JSON(http.StatusOK, i)
}
