package repo

import (
	"github.com/TeeSSK/GoSimpleREST/model"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db}
}

func (userRepo UserRepoImpl) GetUsers() ([]model.User, error) {
	users := []model.User{}
	result := userRepo.DB.Find(&users)
	if result.Error != nil {
		return []model.User{}, result.Error
	}
	return users, nil
}

func (userRepo UserRepoImpl) GetUserById(id uint) (model.User, error) {
	var user model.User
	result := userRepo.DB.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (userRepo UserRepoImpl) CreateUser(request model.CreateUserRequest) (model.User, error) {
	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	result := userRepo.DB.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (userRepo UserRepoImpl) UpdateUser(id uint, request model.UpdateUserRequest) (model.User, error) {
	var user model.User
	user, err := userRepo.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	result := userRepo.DB.Model(&user).Select("Name", "Email").Updates(model.User{Name: request.Name, Email: request.Email})
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (userRepo UserRepoImpl) DeleteUser(id uint) (model.User, error) {
	var user model.User
	result := userRepo.DB.Delete(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
