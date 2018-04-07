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
	router.POST("/profile", profile.RH_Create)
	router.GET("/profile", profile.RH_Find)
	//router.GET("/profile/:id", profile.FindByIdProfileRouteHandler)
	//router.PUT("/profile/:id", profile.UpdateProfileRouteHandler)
	//router.DELETE("/profile/:id", profile.DeleteByIdProfileRouteHandler)

	//: Run router
	API_EXPOSED_PORT := os.Getenv("API_EXPOSED_PORT")
	API_EXPOSED_PORT = fmt.Sprintf(":%s", API_EXPOSED_PORT)
	router.Run(API_EXPOSED_PORT)
}
