package ports

import "coin/graph/model"

type UserService interface {
	AddUser(model.UserInput) (*model.User, error)
}
