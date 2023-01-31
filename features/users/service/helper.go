package service

import (
	"crypto/rand"
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
	"math/big"
)

func createRequestToModel(data dto.UserCreateRequest) domain.User {
	return domain.User{
		Name:     data.Name,
		Nik:      data.Nik,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
	}
}

func modelToResponse(data domain.User) dto.UserResponse {
	return dto.UserResponse{
		Id:        data.Id,
		Name:      data.Name,
		Nik:       data.Nik,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		Role:      data.Role,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func responseToToken(data dto.UserResponse, token string) dto.UserDataToken {
	return dto.UserDataToken{
		Id:    data.Id,
		Role:  data.Role,
		Phone: data.Phone,
		Email: data.Email,
		Token: token,
	}
}

func generateVerCode(length int) (string, error) {
	seed := "012345679"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}
