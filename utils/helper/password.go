package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, cost int) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash)
}
