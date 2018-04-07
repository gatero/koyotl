package profile

import (
	firebaseAdmin "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	googleOption "google.golang.org/api/option"
)

func GetToken(token string) (*auth.Token, error) {
	opt := googleOption.WithCredentialsFile("firebase/config.development.json")
	app, error := firebaseAdmin.NewApp(context.Background(), nil, opt)

	if error != nil {
		return nil, error
	}

	client, error := app.Auth(context.Background())

	if error != nil {
		return nil, error
	}

	return client.VerifyIDToken(token)
}
