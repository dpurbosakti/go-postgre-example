package service

import (
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"
)

func requestToModel(data dto.UserCreateRequest) domain.User {
	return domain.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
	}
}

func modelToResponse(data domain.User) dto.UserCreateResponse {
	return dto.UserCreateResponse{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}
