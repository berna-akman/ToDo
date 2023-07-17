package presentation

import (
	"to-do-api/cb/domain/user"
)

type UserService interface {
	GetAllUsers() (*[]user.User, error)
	CreateUser(user2 user.User) (*user.CreateResponse, error)
}

type userService struct {
	r user.UserRepository
}

func NewUserService(repository user.UserRepository) UserService {
	return userService{repository}
}

func (s userService) GetAllUsers() (*[]user.User, error) {
	users, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s userService) CreateUser(user user.User) (*user.CreateResponse, error) {
	user.CardIDs = []string{}
	return s.r.CreateUser(user)
}
