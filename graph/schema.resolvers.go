package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"coin/graph/generated"
	"coin/graph/model"
	"context"
	"fmt"

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
		fmt.Println(err)
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
