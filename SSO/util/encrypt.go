package util

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ValidatePassword(origin, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(origin))
}
