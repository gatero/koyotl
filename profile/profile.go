package profile

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	fmt.Println("profile test create method")
	c.String(http.StatusOK, "hello")
}

func (m *Model) Find(c *gin.Context) {
	fmt.Println("profile test create method")
}

func (m *Model) Update(c *gin.Context) {
	fmt.Println("profile test create method")
}

func (m *Model) Delete(c *gin.Context) {
	fmt.Println("profile test create method")
}
