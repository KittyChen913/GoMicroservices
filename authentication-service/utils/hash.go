package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hashPassword string, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))
	return err
}
