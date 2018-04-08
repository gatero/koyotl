package profile

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Find method get all the user instances
// that exist into the database
// TODO: Filter the results through a query
func Find(query Profile) ([]Profile, error) {
	// get the collection pointer
	c, _ := Collection()
	// declare and empty array
	var p []Profile
	// try to find all the profiles
	// and store them into the profiles slice
	fmt.Printf("\n\n QUERY: %#v\n\n", query)
	if e := c.Find(query).All(&p); e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		return nil, e
	}
	return p, nil
}

// RH_Find is the route HandlerFunc for find
func RH_Find(c *gin.Context) {
	var form Profile
	c.BindJSON(&form)
	// try to find all the profiles
	// and store them into the profiles slice
	p, e := Find(form)

	if e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		c.String(http.StatusInternalServerError, e.Error())
	}
	// return the query results
	c.JSON(http.StatusOK, p)
}
