package profile

import (
	"app/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Model struct{}

type Profile struct {
	FirstName string
	LastName  string
	Birthday  string
	Email     []string
}

func (m *Model) Create(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func (m *Model) Find(c *gin.Context) {
	mongo := db.Mongo{}
	session, _ := mongo.Open()
	fmt.Printf("%s", session)
	c.String(http.StatusOK, "hello")
}

func (m *Model) Update(c *gin.Context) {
	fmt.Println("profile test create method")
}

func (m *Model) Delete(c *gin.Context) {
	fmt.Println("profile test create method")
}
