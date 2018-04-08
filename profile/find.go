package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RH_Find(ctx *gin.Context) {
	c, _ := ProfileC()

	var profiles []Profile
	if e := c.Find(nil).All(&profiles); e != nil {
		ctx.JSON(http.StatusInternalServerError, e)
	}

	ctx.JSON(http.StatusOK, profiles)
}
