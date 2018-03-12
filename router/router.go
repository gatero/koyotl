package router

import (
	"app/profile"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct{}

func (c Config) Up() {
	router := gin.Default()

	profile := &profile.Model{}

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
