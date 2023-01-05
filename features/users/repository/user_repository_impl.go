package repository

import (
	"fmt"
	"learn-echo/features/users/model/domain"

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
