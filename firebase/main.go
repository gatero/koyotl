package firebase

import (
	"app/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

var TOKEN string = "Token"

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN)

		if len(token) > 0 {
			t, e := profile.GetToken(token)
			profile.VerifyToken(t)

			if e != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"type":    "Token",
					"message": e.Error(),
				})
			}

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"type":    "Token",
				"message": "No token present in the request",
			})
		}

	}
}
