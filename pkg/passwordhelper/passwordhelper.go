package passwordhelper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(input string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func ComparePassword(input string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input))
	return err
}
