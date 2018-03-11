package router

import (
	"app/profile"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DB *gorm.DB
}

func (c Config) Up() {
	router := gin.Default()

	profile := &profile.Model{DB: c.DB}

	router.GET("/profile", profile.Find)
	router.POST("/profile", profile.Create)
	router.PUT("/profile", profile.Update)
	router.DELETE("/profile", profile.Delete)

	API_EXPOSED_PORT := fmt.Sprintf(
		":%s",
		os.Getenv("API_EXPOSED_PORT"),
	)
	router.Run(API_EXPOSED_PORT)
}
