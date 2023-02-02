package repository

import (
	"errors"
	"fmt"
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
	"learn-echo/pkg/pagination"

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

func (repository *UserRepositoryImpl) GetDetail(tx *gorm.DB, userId int) (domain.User, error) {
	var user *domain.User
	result := tx.First(&user, userId)
	if result.Error != nil {
		return domain.User{}, fmt.Errorf("data user dengan user id %d tidak ditemukan", userId)
	}

	return *user, nil
}

func (repository UserRepositoryImpl) GetList(tx *gorm.DB, page pagination.Pagination) (pagination.Pagination, error) {
	var users []domain.User

	tx.Scopes(pagination.Paginate(users, &page, tx)).Find(&users)
	page.Rows = users

	return page, nil
}

func (repository UserRepositoryImpl) Delete(tx *gorm.DB, userId int) error {
	var user domain.User
	result := tx.Delete(&user, userId)
	if result.Error != nil {
		return fmt.Errorf("data user dengan user id %d tidak ditemukan", userId)
	}

	return nil
}

func (repository UserRepositoryImpl) Update(tx *gorm.DB, input domain.User) (domain.User, error) {
	result := tx.Save(&input)
	if result.Error != nil {
		return domain.User{}, errors.New("gagal memperbarui dan menyimpan data user")
	}

	return input, nil
}

func (repository UserRepositoryImpl) CheckDuplicate(tx *gorm.DB, input domain.User) error {
	var count int64
	if resultEmail := tx.Model(&domain.User{}).Where("email = ? ", input.Email).Count(&count); resultEmail.Error != nil {
		return errors.New("error checking email")
	}
	if count > 0 {
		return errors.New("email already exists in database")
	}
	if resultNik := tx.Model(&domain.User{}).Where("nik = ? ", input.Nik).Count(&count); resultNik.Error != nil {
		return errors.New("error checking nik")
	}
	if count > 0 {
		return errors.New("nik already exists in database")
	}
	if resultPhone := tx.Model(&domain.User{}).Where("phone = ? ", input.Phone).Count(&count); resultPhone.Error != nil {
		return errors.New("error checking phone")
	}
	if count > 0 {
		return errors.New("phone already exists in database")
	}

	return nil
}

func (repository UserRepositoryImpl) CheckEmail(tx *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	result := tx.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return domain.User{}, errors.New("error checking email")
	}
	return user, nil
}
