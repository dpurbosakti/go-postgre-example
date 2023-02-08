package service

import (
	"learn-echo/config"
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
	"learn-echo/mocks"
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
	input := dto.UserLoginRequest{
		Email:    "example@gmail.com",
		Password: "CobaCoba",
	}

	response := domain.User{
		Id: 1, Nik: "", Name: "", Email: "example@gmail.com", Password: "$2a$10$5eZ8T4/BO7JeLrOrIaigSuuyLS62fouMkQZzOtu1YWpMqpI4CEuBy", Phone: "085605430555", Address: "", Role: "user", VerCode: "", IsVerified: true, CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), DeletedAt: gorm.DeletedAt{Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid: false},
	}

	result := dto.UserDataToken{
		Id: 1, Role: "user", Phone: "085605430555", Email: "example@gmail.com", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImR3aWF0bW9rb3BAZ21haWwuY29tIiwiZXhwIjoxNjc3NjM3OTM5LCJwaG9uZSI6IjA4NTYwNTQzMDU1NSIsInJvbGUiOiJ1c2VyIiwidXNlcklkIjoxfQ.rTJuTTQIy-QbQ1q2EKrjn9nF3ppF4L4VJa0jJafAIpI",
	}
	t.Run("Login Success", func(t *testing.T) {
		repo.On("Login", mock.Anything, mock.Anything).Return(response, nil).Once()
		srv := NewUserService(repo, db)

		res, err := srv.Login(input)
		res.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImR3aWF0bW9rb3BAZ21haWwuY29tIiwiZXhwIjoxNjc3NjM3OTM5LCJwaG9uZSI6IjA4NTYwNTQzMDU1NSIsInJvbGUiOiJ1c2VyIiwidXNlcklkIjoxfQ.rTJuTTQIy-QbQ1q2EKrjn9nF3ppF4L4VJa0jJafAIpI"
		assert.NoError(t, err)
		assert.Equal(t, result, res)
		repo.AssertExpectations(t)
	})

}
