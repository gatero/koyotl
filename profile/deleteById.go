package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

// RH_Delete method removes a profile instance using the uid pased through
// URL params
func RH_DeleteById(ctx *gin.Context) {
	c, _ := Collection()
	// create id bson object that contains the parsed param id
	id := bson.M{"_id": bson.ObjectIdHex(ctx.Param("id"))}

	if err := c.Remove(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.String(http.StatusOK, "The profile %s was deleted successfully !", ctx.Param("id"))
}
