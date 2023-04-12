package repository

import "challange-10/entity"

type UserRepository interface {
	AddUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindById(id uint) (entity.User, error)
}
