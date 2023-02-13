package dto

type AccountCreateRequest struct {
	Type string `json:"type" validate:"required"`
}
