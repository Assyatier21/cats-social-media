package helper

import (
	"math/rand"
	"time"

	"github.com/backendmagang/project-1/utils/constant"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(constant.CHARSET))
		password[i] = constant.CHARSET[randomIndex]
	}

	return string(password)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
