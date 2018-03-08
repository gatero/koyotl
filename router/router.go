package router

import (
	"app/profile"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DB *gorm.DB
}

func (c Config) Up() {
	router := gin.Default()

	profile := profile.Model{DB: c.DB}
	router.GET("/profile", profile.Find)
}
