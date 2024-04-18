package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/goropencho/golang-gin/data/request"
	"github.com/goropencho/golang-gin/data/response"
	"github.com/goropencho/golang-gin/helper"
	model "github.com/goropencho/golang-gin/model/users"
	repository "github.com/goropencho/golang-gin/repository/users"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UsersResponse {
	result := u.UserRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			Id:      int(value.ID),
			Name:    value.Name,
			Address: value.Address,
			Email:   value.Email,
		}
		users = append(users, user)
	}
	return users
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(userId int) response.UsersResponse {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		Id:      int(userData.ID),
		Name:    (userData.Name),
		Address: (userData.Address),
		Email:   (userData.Email),
	}
	return userResponse
}

// Update implements UserService.
func (u *UserServiceImpl) Update(users request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.Name = users.Name
	userData.Email = users.Email
	userData.Address = users.Address
	u.UserRepository.Update(userData)
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (u UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}
	u.UserRepository.Save(userModel)
}
