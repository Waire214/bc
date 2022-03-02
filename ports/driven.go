package ports
import "coin/graph/model"

type UserRepository interface {
	AddUser(model.UserInput) (*model.User, error)
}