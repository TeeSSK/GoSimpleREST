package services

import (
	"github.com/TeeSSK/GoSimpleREST/model"
	repo "github.com/TeeSSK/GoSimpleREST/repo"
)

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{userRepo}
}

func (userService *UserServiceImpl) GetUsers() ([]model.User, error) {
	res, err := userService.userRepo.GetUsers()
	return res, err
}

func (userService UserServiceImpl) GetUserById(id uint) (model.User, error) {
	res, err := userService.userRepo.GetUserById(id)
	return res, err
}

func (userService UserServiceImpl) CreateUser(request model.CreateUserRequest) (model.User, error) {
	// TODO: add salt and hash password
	res, err := userService.userRepo.CreateUser(request)
	return res, err
}

func (userService UserServiceImpl) UpdateUser(id uint, request model.UpdateUserRequest) (model.User, error) {
	res, err := userService.userRepo.UpdateUser(id, request)
	return res, err
}

func (userService UserServiceImpl) DeleteUser(id uint) (model.User, error) {
	res, err := userService.userRepo.DeleteUser(id)
	return res, err
}
