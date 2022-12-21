package user

import (
	gh "learn-echo/pkg/gormhelper/types"
)

type Core struct {
	ID       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Phone    string `validate:"required"`
	Address  string
	Role     string `validate:"required"`
	gh.DateModel
}
