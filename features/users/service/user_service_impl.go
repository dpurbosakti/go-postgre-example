package service

import (
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
	"learn-echo/features/users/repository"
	"learn-echo/middlewares"
	ph "learn-echo/pkg/passwordhelper"

	"fmt"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Create(input dto.UserCreateRequest) (result dto.UserResponse, err error) {
	hashPassword, errHash := ph.HashPassword(input.Password)
	if errHash != nil {
		return result, errHash
	}
	input.Password = hashPassword
	data := requestToModel(input)
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultData, err := service.UserRepository.Create(tx, data)
		if err != nil {
			return err
		}
		result = modelToResponse(resultData)
		return nil
	})
	if err != nil {
		return dto.UserResponse{}, err
	}

	return result, nil
}

func (service *UserServiceImpl) Login(input dto.UserLoginRequest) (result dto.UserDataToken, err error) {
	var resultData domain.User
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := service.UserRepository.Login(tx, input)
		if err != nil {
			return err
		}
		resultData = resultRepo
		return nil
	})

	if err != nil {
		return result, err
	}
	errCrypt := ph.ComparePassword(resultData.Password, input.Password)
	if errCrypt != nil {
		return result, fmt.Errorf("password incorrect")
	}

	dataToken := modelToResponse(resultData)
	token, err := middlewares.CreateToken(dataToken)
	if err != nil {
		return result, err
	}

	result = responseToToken(dataToken, token)
	return result, nil
}
