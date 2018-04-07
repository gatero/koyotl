package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RH_Find(ctx *gin.Context) {
	c, e := ProfileC()
	if e != nil {
		panic(e)
	}

	var profiles []Profile
	c.Find(nil).All(&profiles)

	ctx.JSON(http.StatusOK, profiles)
}
