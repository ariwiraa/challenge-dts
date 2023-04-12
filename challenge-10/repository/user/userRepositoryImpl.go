package repository

import (
	"challange-10/entity"
	"log"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// FindById implements UserRepository
func (repository *UserRepositoryImpl) FindById(id uint) (entity.User, error) {
	var user entity.User

	err := repository.Db.Debug().Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// FindByEmail implements UserRepository
func (repository *UserRepositoryImpl) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := repository.Db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// AddUser implements UserRepository
func (repository *UserRepositoryImpl) AddUser(user entity.User) (entity.User, error) {
	err := repository.Db.Create(&user).Error
	if err != nil {
		log.Fatal("error creating book data: ", err)
	}
	return user, nil
}
