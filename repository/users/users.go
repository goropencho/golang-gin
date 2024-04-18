package repository

import (
	"errors"

	"github.com/goropencho/golang-gin/data/request"
	"github.com/goropencho/golang-gin/helper"
	"github.com/goropencho/golang-gin/model/users"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(users users.User)
	Update(users users.User)
	Delete(userId int)
	FindById(userId int) (users users.User, err error)
	FindAll() []users.User
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user users.User
	result := u.DB.Where("ID = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []users.User {
	var users []users.User
	u.DB.Find(&users)
	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (users users.User, err error) {
	result := u.DB.Find(&users, userId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user not found")
	}
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(users users.User) {
	result := u.DB.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(users users.User) {
	var updateUser = request.UpdateUserRequest{
		Id:      int(users.ID),
		Name:    users.Name,
		Address: users.Address,
	}
	result := u.DB.Model(&users).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func NewUsersRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}
