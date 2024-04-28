package helper

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestGenerateUUIDString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Generated UUID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUUIDString()

			_, err := uuid.Parse(got)
			if err != nil {
				t.Errorf("GenerateUUIDString() returned an invalid UUID: %v", err)
			}

			if strings.Contains(got, "-") {
				t.Errorf("GenerateUUIDString() generated a UUID with dashes: %s", got)
			}
		})
	}
}
