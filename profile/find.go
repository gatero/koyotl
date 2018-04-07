package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(p []Profile) []Profile {
	if c, e := ProfileC(); e == nil {
		c.Find(nil).All(&p)
	}
	return p
}

func RH_Find(c *gin.Context) {
	var profiles []Profile
	Find(profiles)
	c.JSON(http.StatusOK, profiles)
}
