package service

import (
	"errors"
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
	"learn-echo/features/users/repository"
	"learn-echo/middlewares"
	eh "learn-echo/pkg/emailhelper"
	"learn-echo/pkg/pagination"
	ph "learn-echo/pkg/passwordhelper"

	"github.com/jinzhu/copier"

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
	// var errChan = make(chan error)
	hashPassword, errHash := ph.HashPassword(input.Password)
	if errHash != nil {
		return result, errHash
	}
	input.Password = hashPassword
	data := createRequestToModel(input)
	verCode, _ := generateVerCode(6)
	data.IsVerified = false
	data.VerCode = verCode
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := service.UserRepository.Create(tx, data)
		if err != nil {
			return err
		}
		result = modelToResponse(resultRepo)
		return nil
	})
	if err != nil {
		return dto.UserResponse{}, err
	}
	err = eh.SendEmail(data)
	if err != nil {
		return dto.UserResponse{}, errors.New("failed to send email verification code: " + err.Error())
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

func (service *UserServiceImpl) GetDetail(userId int) (result dto.UserResponse, err error) {
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := service.UserRepository.GetDetail(tx, userId)
		if err != nil {
			return err
		}
		result = modelToResponse(resultRepo)
		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *UserServiceImpl) GetList(page pagination.Pagination) (result pagination.Pagination, err error) {
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := service.UserRepository.GetList(tx, page)
		if err != nil {
			return err
		}
		result = resultRepo
		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *UserServiceImpl) Delete(userId int) (err error) {
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		err := service.UserRepository.Delete(tx, userId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) Update(input dto.UserUpdateRequest, userId int) (result dto.UserResponse, err error) {
	if input.Password != nil {
		hashPassword, errHash := ph.HashPassword(*input.Password)
		if errHash != nil {
			return result, errHash
		}
		input.Password = &hashPassword
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultGet, err := service.UserRepository.GetDetail(tx, userId)
		if err != nil {
			return err
		}
		fmt.Println(resultGet.Id)
		err = copier.Copy(&resultGet, &input)
		if err != nil {
			return err
		}
		fmt.Println(resultGet.Id)
		resultUpdate, err := service.UserRepository.Update(tx, resultGet)
		if err != nil {
			return err
		}
		result = modelToResponse(resultUpdate)
		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}
