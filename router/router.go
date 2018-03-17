package router

import (
	"app/cors"
	"app/profile"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct{}

func (c Config) Up() {
	router := gin.Default()
	router.Use(cors.Setup())

	//: Routes
	router.GET("/profile", profile.Find)
	router.GET("/profile/:id", profile.FindById)
	router.POST("/profile", profile.Create)
	router.PUT("/profile", profile.Update)
	router.DELETE("/profile/:id", profile.DeleteById)

	//: Run router
	API_EXPOSED_PORT := os.Getenv("API_EXPOSED_PORT")
	API_EXPOSED_PORT = fmt.Sprintf(":%s", API_EXPOSED_PORT)
	router.Run(API_EXPOSED_PORT)
}
