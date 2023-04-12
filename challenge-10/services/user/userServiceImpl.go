package service

import (
	"challange-10/entity"
	"challange-10/helpers"
	repository "challange-10/repository/user"
	"errors"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

// GetUserById implements UserService
func (service *UserServiceImpl) GetUserById(id uint) (entity.User, error) {
	user, err := service.UserRepository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}
	return user, nil
}

// Login implements UserService
func (service *UserServiceImpl) Login(payload entity.UserLogin) (entity.User, error) {
	email := payload.Email
	password := payload.Password

	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	comparePassword := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePassword {
		return user, errors.New("invalid password")
	}

	return user, nil
}

// AddUser implements UserService
func (service *UserServiceImpl) AddUser(request entity.UserRequest) (entity.User, error) {
	user := entity.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: request.Password,
		IsAdmin:  request.IsAdmin,
	}

	newUser, err := service.UserRepository.AddUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}
