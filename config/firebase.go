package config

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	firebaseApp *firebase.App
	authClient  *auth.Client
)

func ConnectFirebase() *firebase.App {

	LoadEnv()

	firebaseURL := os.Getenv("FIREBASE_URL")

	opt := option.WithCredentialsFile("credential.json")
	config := &firebase.Config{
		DatabaseURL: firebaseURL,
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	firebaseApp = app

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth: %v", err)
	}

	authClient = client

	fmt.Println("Connected to Firebase")

	return app
}
