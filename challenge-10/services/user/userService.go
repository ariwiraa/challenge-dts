package service

import "challange-10/entity"

type UserService interface {
	AddUser(request entity.UserRequest) (entity.User, error)
	Login(payload entity.UserLogin) (entity.User, error)
	GetUserById(id uint) (entity.User, error)
}
