package services

import "github.com/TeeSSK/GoSimpleREST/model"

type UserService interface {
	GetUsers() ([]model.User, error)
	GetUserById(id uint) (model.User, error)
	CreateUser(request model.CreateUserRequest) (model.User, error)
	UpdateUser(id uint, request model.UpdateUserRequest) (model.User, error)
	DeleteUser(id uint) (model.User, error)
}
