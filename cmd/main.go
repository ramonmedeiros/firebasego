package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	fmt.Println("running")

	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
}
