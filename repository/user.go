package repository

import (
	"coin/graph/model"
	"coin/helper"
	"coin/ports"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx    = context.TODO()
	// source string
	// target string
)

type userInfra struct {
	UserCollection *mongo.Collection
}

func NewUserRepository(conn *mongo.Database) ports.UserRepository {
	return &userInfra{
		UserCollection: conn.Collection("user"),
	}
}

func (collection *userInfra) AddUser(user model.UserInput) (*model.User, error) {
	helper.LogEvent("INFO", "Persisting new user")

	_, err := collection.UserCollection.InsertOne(
		ctx,
		user,
	)
	if err != nil {
		return &model.User{}, helper.ErrorMessage(helper.MongoDBError, err.Error())
	}

	helper.LogEvent("INFO", "Persisting new user successful")
	newUser := model.User(user)
	return &newUser, nil
}
