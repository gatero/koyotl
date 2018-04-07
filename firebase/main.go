package firebase

import (
	"app/profile"
	"fmt"
	"net/http"

	firebaseAdmin "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	googleOption "google.golang.org/api/option"
)

var TOKEN string = "Token"

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN)
		fmt.Printf("\n\nToken header %s\n\n", token)

		if len(token) > 0 {
			opt := googleOption.WithCredentialsFile("firebase/config.development.json")
			app, error := firebaseAdmin.NewApp(context.Background(), nil, opt)

			if error != nil {
				panic(error)
			}

			client, error := app.Auth(context.Background())
			fmt.Printf("\n\nClient: %s\n\n", client)

			if error != nil {
				panic(error)
			}

			decodedToken, error := client.VerifyIDToken(token)
			fmt.Printf("\n\n DECODED TOKEN: %s\n\n", decodedToken)
			profile.VerifyToken(decodedToken)

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
				"type":    "Token",
				"message": "No token present in the request",
			})
		}

	}
}
