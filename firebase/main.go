package firebase

import (
	"app/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This is the label of Token header
var TOKEN string = "Token"

// VerifyToken is a middleware to check the token status,
// if it expires then this method returns an unauthorized status code(401)
// else if there isn't token on the request then also
// return the corresponding unauthorized status code(401)
func VerifyToken() gin.HandlerFunc {
	// get the gin context
	return func(ctx *gin.Context) {
		// get the Token header
		token := ctx.GetHeader(TOKEN)
		// Verify the existance of token header
		if len(token) > 0 {
			// process the token and get the status
			t, e := profile.GetToken(token)
			// if token is invalid or expired then returns
			// the corresponding status code(401)
			if e != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"type":    "Token",
					"message": e.Error(),
				})
			}
			// Verify if the profile exists
			profile.VerifyProfile(t)
			// continue
			ctx.Next()
		} else {
			// if the request hasn't Token header
			// return the corresponding unauthorized status code
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"type":    "Token",
				"message": "No token present in the request",
			})
		}
	}
}
