package profile

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	mgo "gopkg.in/mgo.v2"
)

type Model struct {
	DB *gorm.DB
}

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
	session, error := mgo.Dial(os.Getenv("MONGO_CONTAINER"))
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s", session)
	c.String(http.StatusOK, "hello")
}

func (m *Model) Update(c *gin.Context) {
	fmt.Println("profile test create method")
}

func (m *Model) Delete(c *gin.Context) {
	fmt.Println("profile test create method")
}
