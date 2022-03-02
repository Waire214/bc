package graph

import (
	"coin/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct {
// 	db *mongo.Database
// }

// func NewResolverHandler(db *mongo.Database) *Resolver {
// 	return &Resolver{
// 		db: db,
// 	}
// }

type Resolver struct {
	userService ports.UserService
}

func NewResolverHandler(userService ports.UserService) *Resolver {
	return &Resolver{
		userService: userService,
	}
}
