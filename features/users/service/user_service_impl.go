package service

import (
	"learn-echo/features/users/model/dto"
	"learn-echo/features/users/repository"
	ph "learn-echo/pkg/passwordhelper"

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

func (service *UserServiceImpl) Create(input dto.UserCreateRequest) (result dto.UserCreateResponse, err error) {
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
		return dto.UserCreateResponse{}, err
	}

	return result, nil
}
