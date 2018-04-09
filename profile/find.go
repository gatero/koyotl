package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Find method get all the user instances
// that exist into the database
func Find(query map[string]interface{}) ([]Profile, error) {
	// get the collection pointer
	c, _ := Collection()
	// declare and empty array
	var p []Profile
	// try to find all the profiles
	// and store them into the profiles slice
	if len(query) == 0 {
		query = nil
	}

	if e := c.Find(query).All(&p); e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		return nil, e
	}
	return p, nil
}

// RH_Find is the route HandlerFunc for find
func RH_Find(c *gin.Context) {
	var query map[string]interface{}
	c.ShouldBindJSON(&query)
	// try to find all the profiles
	// and store them into the profiles slice
	p, e := Find(query)

	if e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		c.String(http.StatusInternalServerError, e.Error())
	}
	// return the query results
	c.JSON(http.StatusOK, p)
}
