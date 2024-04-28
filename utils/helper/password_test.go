package helper

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomPassword(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name      string
		args      args
		wantRegex string
	}{
		{
			name:      "Length 8",
			args:      args{length: 8},
			wantRegex: "^[A-Za-z0-9]{8}$",
		},
		{
			name:      "Length 12",
			args:      args{length: 12},
			wantRegex: "^[A-Za-z0-9]{12}$",
		},
		{
			name:      "Length 16",
			args:      args{length: 16},
			wantRegex: "^[A-Za-z0-9]{16}$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomPassword(tt.args.length)
			match, _ := regexp.MatchString(tt.wantRegex, got)

			assert.True(t, match, "Generated password does not match expected regex pattern")
			assert.Equal(t, tt.args.length, len(got), "Generated password length does not match expected length")
		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{password: "password123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) != tt.wantErr || (tt.wantErr && got != "") {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
