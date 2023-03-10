package service

import (
	"learn-echo/features/accounts/models/domain"
	"learn-echo/features/accounts/models/dto"
)

func modelToResponse(input domain.Account) dto.AccountResponse {
	return dto.AccountResponse{
		Id:        input.Id,
		Name:      input.Name,
		Type:      input.Type,
		Balance:   input.Balance,
		User_Id:   input.User_Id,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		DeletedAt: input.DeletedAt,
	}
}
