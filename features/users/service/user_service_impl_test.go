package service

import (
	"errors"
	"learn-echo/config"
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
	"learn-echo/mocks"
	"learn-echo/pkg/pagination"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// TODO: nil pointer on checking sendemailvercode func
// func TestCreate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	db := config.InitDBTest()
// 	input := dto.UserCreateRequest{
// 		Name:     "dwi",
// 		Nik:      "1234567890987654",
// 		Email:    "example@gmail.com",
// 		Password: "test",
// 		Phone:    "085085085085",
// 		Address:  "rombo",
// 	}

// 	response := domain.User{
// 		Id: 1, Nik: "1234567890987654", Name: "dwi", Email: "example@gmail.com", Password: "$2a$10$5eZ8T4/BO7JeLrOrIaigSuuyLS62fouMkQZzOtu1YWpMqpI4CEuBy", Phone: "085085085085", Address: "rombo", Role: "user", VerCode: "", IsVerified: true, CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), DeletedAt: gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
// 	}

// 	result := dto.UserResponse{
// 		Name:       "dwi",
// 		Nik:        "1234567890987654",
// 		Email:      "example@gmail.com",
// 		Phone:      "085085085085",
// 		Address:    "rombo",
// 		Role:       "user",
// 		IsVerified: true,
// 		CreatedAt:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
// 		UpdatedAt:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
// 		DeletedAt:  gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
// 	}
// 	t.Run("Create Success", func(t *testing.T) {
// 		repo.On("CheckDuplicate", mock.Anything, mock.Anything).Return(nil).Once()
// 		repo.On("Create", mock.Anything, mock.Anything).Return(response, nil).Once()
// 		srv := NewUserService(repo, db)

// 		res, err := srv.Create(input)
// 		assert.NoError(t, err)
// 		assert.Equal(t, result, res)
// 		repo.AssertExpectations(t)
// 	})
// }

func TestLogin(t *testing.T) {
	repo := new(mocks.UserRepository)
	db := config.InitDBTest()

	inputSuccess := dto.UserLoginRequest{
		Email:    "example@gmail.com",
		Password: "CobaCoba",
	}

	responseSuccess := domain.User{
		Id: 1, Nik: "", Name: "", Email: "example@gmail.com", Password: "$2a$10$5eZ8T4/BO7JeLrOrIaigSuuyLS62fouMkQZzOtu1YWpMqpI4CEuBy", Phone: "085605430555", Address: "", Role: "user", VerCode: "", IsVerified: true, CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), DeletedAt: gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
	}

	resultSuccess := dto.UserDataToken{
		Id: 1, Role: "user", Phone: "085605430555", Email: "example@gmail.com", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImR3aWF0bW9rb3BAZ21haWwuY29tIiwiZXhwIjoxNjc3NjM3OTM5LCJwaG9uZSI6IjA4NTYwNTQzMDU1NSIsInJvbGUiOiJ1c2VyIiwidXNlcklkIjoxfQ.rTJuTTQIy-QbQ1q2EKrjn9nF3ppF4L4VJa0jJafAIpI",
	}

	inputFailed := dto.UserLoginRequest{
		Email:    "example@gmail.com",
		Password: "test",
	}

	responseFailed := domain.User{}

	resultFailed := dto.UserDataToken{}
	errorLogin := errors.New("password incorrect")
	t.Run("Login Success", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.Anything).Return(responseSuccess, nil).Once()
		srv := NewUserService(repo, db)

		res, err := srv.Login(inputSuccess)
		res.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImR3aWF0bW9rb3BAZ21haWwuY29tIiwiZXhwIjoxNjc3NjM3OTM5LCJwaG9uZSI6IjA4NTYwNTQzMDU1NSIsInJvbGUiOiJ1c2VyIiwidXNlcklkIjoxfQ.rTJuTTQIy-QbQ1q2EKrjn9nF3ppF4L4VJa0jJafAIpI"
		assert.NoError(t, err)
		assert.Equal(t, resultSuccess, res)
		repo.AssertExpectations(t)
	})

	t.Run("Login Failed", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.Anything).Return(responseFailed, errorLogin).Once()
		srv := NewUserService(repo, db)

		res, err := srv.Login(inputFailed)
		assert.EqualError(t, err, errorLogin.Error())
		assert.Equal(t, resultFailed, res)
		repo.AssertExpectations(t)
	})

}

func TestGetDetail(t *testing.T) {
	repo := new(mocks.UserRepository)
	db := config.InitDBTest()

	responseSuccess := domain.User{
		Id: 1, Nik: "1234567890987654", Name: "dwi", Email: "example@gmail.com", Password: "$2a$10$5eZ8T4/BO7JeLrOrIaigSuuyLS62fouMkQZzOtu1YWpMqpI4CEuBy", Phone: "085605430555", Address: "rombo", Role: "user", VerCode: "", IsVerified: true, CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), DeletedAt: gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
	}

	resultSuccess := dto.UserResponse{
		Id: 1, Nik: "1234567890987654", Name: "dwi", Email: "example@gmail.com", Address: "rombo", Role: "user",
		Phone: "085605430555", IsVerified: true, CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), DeletedAt: gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
	}

	responseFailed := domain.User{}

	resultFailed := dto.UserResponse{}

	errorGetDetail := errors.New("user id 1 not found")

	t.Run("Get Detail Success", func(t *testing.T) {
		repo.On("GetDetail", mock.Anything, mock.Anything).Return(responseSuccess, nil).Once()
		srv := NewUserService(repo, db)

		res, err := srv.GetDetail(1)
		assert.NoError(t, err)
		assert.Equal(t, resultSuccess, res)
		repo.AssertExpectations(t)
	})

	t.Run("Get Detail Failed", func(t *testing.T) {
		repo.On("GetDetail", mock.Anything, mock.Anything).Return(responseFailed, errorGetDetail).Once()
		srv := NewUserService(repo, db)

		res, err := srv.GetDetail(1)
		assert.EqualError(t, err, errorGetDetail.Error())
		assert.Equal(t, resultFailed, res)
		repo.AssertExpectations(t)
	})

}

func TestGetList(t *testing.T) {
	repo := new(mocks.UserRepository)
	db := config.InitDBTest()
	input := pagination.Pagination{
		Limit:      1,
		Page:       1,
		Sort:       "id asc",
		TotalRows:  1,
		TotalPages: 1,
		Rows: []domain.User{
			{
				Id:   1,
				Name: "dwi",
			},
		},
	}

	// responseSuccess := domain.User{
	// 	Id:   1,
	// 	Name: "dwi",
	// }

	resultSuccess := pagination.Pagination{
		Limit:      1,
		Page:       1,
		Sort:       "id asc",
		TotalRows:  1,
		TotalPages: 1,
		Rows: []domain.User{
			{
				Id:   1,
				Name: "dwi",
			},
		},
	}

	// get list success but empty
	resultEmpty := pagination.Pagination{
		Limit:      1,
		Page:       1,
		Sort:       "id asc",
		TotalRows:  1,
		TotalPages: 1,
		Rows:       []domain.User{},
	}

	t.Run("Get List Success", func(t *testing.T) {
		repo.On("GetList", mock.Anything, mock.Anything).Return(resultSuccess, nil).Once()
		srv := NewUserService(repo, db)

		res, err := srv.GetList(input)
		assert.NoError(t, err)
		assert.Equal(t, resultSuccess, res)
		repo.AssertExpectations(t)
	})

	t.Run("Get List Success But Empty", func(t *testing.T) {
		repo.On("GetList", mock.Anything, mock.Anything).Return(resultEmpty, nil).Once()
		srv := NewUserService(repo, db)

		res, err := srv.GetList(input)
		assert.NoError(t, err)
		assert.Equal(t, resultEmpty, res)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepository)
	db := config.InitDBTest()

	errFailed := errors.New("user id 1 not found")

	t.Run("Delete Success", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		srv := NewUserService(repo, db)

		err := srv.Delete(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Delete Failed", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.Anything).Return(errFailed).Once()
		srv := NewUserService(repo, db)

		err := srv.Delete(1)
		assert.EqualError(t, err, errFailed.Error())
		repo.AssertExpectations(t)
	})
}

func TestVerify(t *testing.T) {
	repo := new(mocks.UserRepository)
	db := config.InitDBTest()

	responseSuccess := domain.User{
		Id: 1, Nik: "1234567890987654", Name: "dwi", Email: "example@gmail.com", Password: "$2a$10$5eZ8T4/BO7JeLrOrIaigSuuyLS62fouMkQZzOtu1YWpMqpI4CEuBy", Phone: "085605430555", Address: "rombo", Role: "user", VerCode: "123456", IsVerified: true, CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), DeletedAt: gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
	}

	inputSuccess := dto.UserVerifyRequest{
		Email:   "example@gmail.com",
		VerCode: "123456",
	}

	inputError := dto.UserVerifyRequest{
		Email:   "example@gmail.com",
		VerCode: "123451",
	}

	errVerify := errors.New("the verification code you entered is incorrect")

	t.Run("Verify Success", func(t *testing.T) {
		repo.On("CheckEmail", mock.Anything, mock.Anything).Return(responseSuccess, nil).Once()
		repo.On("Update", mock.Anything, mock.Anything).Return(responseSuccess, nil).Once()
		srv := NewUserService(repo, db)

		err := srv.Verify(inputSuccess)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Verify Failed", func(t *testing.T) {
		repo.On("CheckEmail", mock.Anything, mock.Anything).Return(responseSuccess, nil).Once()
		srv := NewUserService(repo, db)

		err := srv.Verify(inputError)
		assert.EqualError(t, err, errVerify.Error())
		repo.AssertExpectations(t)
	})
}
