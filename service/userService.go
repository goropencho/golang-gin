package service

import (
	"github.com/goropencho/golang-gin/data/request"
	"github.com/goropencho/golang-gin/data/response"
)

type UserService interface {
	Create(users request.CreateUserRequest)
	Update(users request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
