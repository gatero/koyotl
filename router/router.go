package router

import (
	"app/cors"
	"app/firebase"
	"app/profile"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	router.Use(cors.Setup())
	router.Use(firebase.VerifyToken())

	//: Profile routes
	router.POST("/profile", profile.Create)
	router.GET("/profile", profile.Find)
	router.GET("/profile/:id", profile.FindById)
	router.PUT("/profile/:id", profile.Update)
	router.DELETE("/profile/:id", profile.DeleteById)

	//: Run router
	API_EXPOSED_PORT := os.Getenv("API_EXPOSED_PORT")
	API_EXPOSED_PORT = fmt.Sprintf(":%s", API_EXPOSED_PORT)
	router.Run(API_EXPOSED_PORT)
}
