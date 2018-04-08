package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func RH_FindById(ctx *gin.Context) {
	collection, _ := ProfileC()
	id := bson.M{"_id": bson.ObjectIdHex(ctx.Param("id"))}
	profile := Profile{}

	if err := collection.Find(id).One(&profile); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, profile)
}
