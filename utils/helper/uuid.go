package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUIDString() string {
	id := uuid.New()
	stringID := id.String()
	uuid := strings.Replace(stringID, "-", "", -1)
	return uuid
}
