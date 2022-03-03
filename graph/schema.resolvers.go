package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"coin/graph/generated"
	"coin/graph/model"
	"context"
	"log"

	"github.com/google/uuid"
)

func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	genReference := uuid.New().String()
	user := model.UserInput{
		ID:                genReference,
		Name:              input.Name,
		BankName:          input.BankName,
		BankCode:          input.BankCode,
		BankAccountNumber: input.BankAccountNumber,
	}
	result, err := r.userService.AddUser(user)
	if err != nil {
		log.Println(err)
		return &model.User{}, err
	}

	return result, nil
}

func (r *queryResolver) BankName(ctx context.Context, bankAcountNumber string, bankCode string) (string, error) {
	result, err := r.userService.GetUser(bankAcountNumber, bankCode)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, bankAcountNumber string, bankCode string) (string, error) {
	result, err := r.userService.GetUser(bankAcountNumber, bankCode)
	if err != nil {
		return "", err
	}
	return result, nil
}
