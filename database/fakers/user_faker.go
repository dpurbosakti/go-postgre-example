package fakers

import (
	"crypto/rand"
	"learn-echo/features/users/models/domain"
	"math/big"
	"time"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func generateNumber(length int) string {
	seed := "012345679"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return ""
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice)
}

func UserFaker(db *gorm.DB) (result []domain.User) {
	for i := 1; i <= 20; i++ {
		tmp := domain.User{
			Id:         uint(i),
			Nik:        generateNumber(16),
			Name:       faker.Name(),
			Email:      faker.Email(),
			Password:   "$2a$10$5eZ8T4/BO7JeLrOrIaigSuuyLS62fouMkQZzOtu1YWpMqpI4CEuBy",
			Phone:      faker.Phonenumber(),
			Address:    faker.Word(),
			Role:       "user",
			VerCode:    "123456",
			IsVerified: true,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		}
		result = append(result, tmp)
	}
	return result
}
