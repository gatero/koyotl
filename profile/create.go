package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create is a ethod to create user instances
// it recives a struct of Profile type
func Create(p *Profile) (*Profile, error) {
	// get the collection pointer
	c, _ := Collection()

	// try to Insert the profile instance
	if e := c.Insert(p); e != nil {
		// return error
		return nil, e
	}

	// return profile instance
	return p, nil
}

// RH_Create is route HandlerFunc
// to create profile instances through api endpoint
// using the Create method
func RH_Create(ctx *gin.Context) {
	var p Profile
	// binding of profile instance
	if e := ctx.ShouldBind(&p); e == nil {
		// try to create the profile instance using Create method
		if profile, e := Create(&p); e == nil {
			ctx.JSON(http.StatusOK, profile)
		} else {
			ctx.JSON(http.StatusInternalServerError, e)
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, e)
	}
}
