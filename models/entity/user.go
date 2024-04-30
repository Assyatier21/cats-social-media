package entity

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	User struct {
		ID        int          `json:"id" db:"id"`
		Email     string       `json:"email" db:"email"`
		Name      string       `json:"name" db:"name"`
		Password  string       `json:"password" db:"password"`
		CreatedAt time.Time    `json:"created_at" db:"created_at"`
		UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
		DeletedAt sql.NullTime `json:"deleted_at,omitempty" db:"deleted_at"`
	}

	CreateUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Name     string `json:"name" validate:"required,min=5,max=50"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	UserJWT struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Token string `json:"accessToken"`
	}

	UserClaims struct {
		Name      string               `json:"name"`
		Email     string               `json:"email"`
		ExpiredAt time.Time            `json:"expired_at"`
		Claims    jwt.RegisteredClaims `json:"claims"`
	}

	UserClaimsResponse struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)
