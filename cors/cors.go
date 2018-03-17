package cors

import (
	corsHandler "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() gin.HandlerFunc {
	return corsHandler.New(
		corsHandler.Config{
			AllowOrigins: []string{"http://localhost"},
			AllowHeaders: []string{"Authorization"},
		},
	)
}
