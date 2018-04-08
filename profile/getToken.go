package profile

import (
	firebaseAdmin "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	googleOption "google.golang.org/api/option"
)

// GetToken recives a token string and return a decoded token
// or an error depends on the firebase response
func GetToken(token string) (*auth.Token, error) {
	// process the configuration file
	opt := googleOption.WithCredentialsFile("firebase/config.development.json")
	// get the firebae app context
	app, error := firebaseAdmin.NewApp(context.Background(), nil, opt)

	if error != nil {
		return nil, error
	}

	// Get the client context
	client, error := app.Auth(context.Background())

	if error != nil {
		return nil, error
	}

	// return the decoded token or error
	return client.VerifyIDToken(token)
}
