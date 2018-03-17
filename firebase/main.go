package firebase

import (
	"net/http"

	firebaseAdmin "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	googleOption "google.golang.org/api/option"
)

func VerifyAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if len(authorization) > 0 {
			opt := googleOption.WithCredentialsFile("config/firebase-admin-development.json")
			app, error := firebaseAdmin.NewApp(context.Background(), nil, opt)

			if error != nil {
				panic(error)
			}

			client, error := app.Auth(context.Background())

			if error != nil {
				panic(error)
			}

			_, error = client.VerifyIDToken(authorization)

			if error != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"type":    "Token",
					"message": error.Error(),
				})
			}

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"type":    "NoToken",
				"message": "No token present in the request",
			})
		}

	}
}
