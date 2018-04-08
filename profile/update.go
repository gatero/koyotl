package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bson "gopkg.in/mgo.v2/bson"
)

func RH_Update(ctx *gin.Context) {
	c, _ := Collection()
	p := Profile{
		FirstName: ctx.PostForm("firstName"),
		LastName:  ctx.PostForm("lastName"),
		Birthday:  ctx.PostForm("birthday"),
	}
	id := bson.M{"_id": bson.ObjectIdHex(ctx.Param("id"))}

	if e := c.Update(id, &p); e != nil {
		ctx.JSON(http.StatusInternalServerError, e)
	}

	ctx.JSON(http.StatusOK, p)
}
