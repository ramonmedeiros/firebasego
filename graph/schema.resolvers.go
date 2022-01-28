package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
	"github.com/ramonmedeiros/firebasego/graph/generated"
	"github.com/ramonmedeiros/firebasego/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	params := (&auth.UserToCreate{}).
		Email(input.Email).
		EmailVerified(true).
		PhoneNumber(input.Phone).
		Password(input.Password).
		DisplayName(input.DisplayName).
		PhotoURL(input.PhotoURL).
		Disabled(false)
	u, err := r.firebaseClient.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
		return nil, err
	}
	log.Printf("Successfully created user: %v\n", u)

	return &model.User{
		Email:       input.Email,
		Phone:       input.Phone,
		DisplayName: input.DisplayName,
		PhotoURL:    input.PhotoURL,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, email string) (*model.User, error) {
	record, err := r.firebaseClient.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("error while retrieving user: %v\n", err)
		return nil, err
	}

	return &model.User{
		Email:       record.Email,
		Phone:       record.PhoneNumber,
		DisplayName: record.DisplayName,
		PhotoURL:    record.PhotoURL,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
