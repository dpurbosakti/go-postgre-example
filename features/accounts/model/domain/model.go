package domain

import (
	m "learn-echo/features/users/model/domain"
)

type Account struct {
	Id      uint    `json:"id" gorm:"primaryKey"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Balance uint    `json:"balance"`
	User_Id uint    `json:"user_id"`
	User    *m.User `json:"user,omitempty" gorm:"foreignKey:User_Id"`
}
