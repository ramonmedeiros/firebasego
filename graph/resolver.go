package graph

import (
	"firebase.google.com/go/auth"
)

type Resolver struct {
	firebaseClient *auth.Client
}
