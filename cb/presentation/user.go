package presentation

import (
	"to-do-api/cb/domain/user"
	"to-do-api/internal/errors"
)

type UserService interface {
	GetAllUsers() (*[]user.User, error)
	CreateUser(user.User) (*user.CreateResponse, error)
	GetUserByID(string) (*user.User, error)
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

func (s userService) GetUserByID(id string) (*user.User, error) {
	b, err := s.r.GetByID(id)
	if err != nil {
		return nil, errors.New(404, "user not found", "user not found")
	}

	return b, nil
}
