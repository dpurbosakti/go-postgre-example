package service

import (
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
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

// func updateRequestToModel(data dto.UserUpdateRequest) domain.User {
// 	return domain.User{
// 		Name:     *data.Name,
// 		Password: *data.Password,
// 		Phone:    *data.Phone,
// 		Address:  *data.Address,
// 	}
// }
