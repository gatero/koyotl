package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(p *Profile) (*Profile, error) {
	c, e := ProfileC()
	if e != nil {
		panic(e)
	}

	if e := c.Insert(p); e != nil {
		return nil, e
	}

	return p, nil
}

func RH_Create(c *gin.Context) {
	var p Profile
	if e := c.ShouldBind(&p); e == nil {
		if profile, e := Create(&p); e == nil {
			c.JSON(http.StatusOK, profile)
		}
	}
}
