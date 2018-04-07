package cors

import (
	"fmt"
	"os"

	corsHandler "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() gin.HandlerFunc {
	STATIC_CONTAINER := os.Getenv("STATIC_CONTAINER")
	STATIC_CONTAINER = fmt.Sprintf("https://%s", STATIC_CONTAINER)

	return corsHandler.New(
		corsHandler.Config{
			AllowOrigins: []string{STATIC_CONTAINER},
			AllowHeaders: []string{"Token"},
		},
	)
}
