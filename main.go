package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("could not get home dir")
	}

	opt := option.WithCredentialsFile(homeDir + "/.config/firebase_admin_sa.json")
	conf := &firebase.Config{ProjectID: "booming-cairn-261420"}

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Println("error initializing app")
	}

	// start graphql server
	client, err := app.Auth(context.Background())
	if err != nil {
		panic("could not start firebase auth")
	}
	startServer(client)
}
