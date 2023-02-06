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
	"strings"

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
		err := service.UserRepository.CheckDuplicate(tx, data)
		if err != nil {
			return err
		}
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
	err = eh.SendEmailVerCode(data)
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

	if !resultData.IsVerified {
		return result, errors.New("account is unverified")
	}
	errCrypt := ph.ComparePassword(resultData.Password, input.Password)
	if errCrypt != nil {
		return result, errors.New("password incorrect")
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
	if page.Sort != "" {
		tmp := strings.Replace(page.Sort, "_", " ", 1)
		page.Sort = tmp
	}
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
		err = copier.Copy(&resultGet, &input)
		if err != nil {
			return err
		}
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

func (service *UserServiceImpl) Verify(input dto.UserVerifyRequest) (err error) {
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		result, err := service.UserRepository.CheckEmail(tx, input.Email)
		if err != nil {
			return err
		}
		if result.VerCode != input.VerCode {
			return errors.New("the verification code you entered is incorrect")
		}

		result.IsVerified = true
		_, errSave := service.UserRepository.Update(tx, result)
		if errSave != nil {
			return errSave
		}

		return nil
	})
	if err != nil {
		fmt.Println("error", err.Error())
		return err
	}

	return nil
}

func (service *UserServiceImpl) RefreshVerCode(input dto.UserVerCodeRequest) (err error) {
	var dataUser domain.User
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := service.UserRepository.CheckEmail(tx, input.Email)
		if err != nil {
			return err
		}
		if resultRepo.IsVerified {
			return errors.New("account is verified no need to refresh/change verification code")
		}
		verCode, _ := generateVerCode(6)
		resultRepo.VerCode = verCode
		_, err = service.UserRepository.Update(tx, resultRepo)
		if err != nil {
			return err
		}
		dataUser = resultRepo
		return nil
	})
	if err != nil {
		return err
	}

	err = eh.SendEmailVerCode(dataUser)
	if err != nil {
		return errors.New("failed to send email verification code: " + err.Error())
	}
	return nil
}
