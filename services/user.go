package services

import (
	"coin/graph/model"
	"coin/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (serve *userService) AddUser(user model.UserInput) (*model.User, error) {
	return serve.userRepository.AddUser(user)
}

func (serve *userService) GetUser(accountNumber string, bankCode string) (string, error) {
	return serve.userRepository.GetUser(accountNumber, bankCode)
}
