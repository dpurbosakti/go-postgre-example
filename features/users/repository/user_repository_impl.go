package repository

import (
	"fmt"
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"

	"gorm.io/gorm"
)

const source = "user"

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(tx *gorm.DB, input domain.User) (domain.User, error) {
	result := tx.Create(&input)
	if result.Error != nil {
		return domain.User{}, fmt.Errorf("failed to create data %s", source)
	}

	return input, nil
}

func (repository *UserRepositoryImpl) Login(tx *gorm.DB, input dto.UserLoginRequest) (domain.User, error) {
	var user domain.User
	result := tx.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return domain.User{}, fmt.Errorf("data user dengan email %s tidak ditemukan", input.Email)
	}

	return user, nil
}
