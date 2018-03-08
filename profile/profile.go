package profile

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Model struct {
	DB *gorm.DB
}

func (m *Model) Create(c *gin.Context) {
	fmt.Println("profile test create method")
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
