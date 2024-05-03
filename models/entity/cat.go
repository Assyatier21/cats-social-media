package entity

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type (
	Cat struct {
		ID               int            `json:"id" db:"id"`
		UserID           int            `json:"user_id" db:"user_id"`
		Name             string         `json:"name" db:"name"`
		Race             string         `json:"race" db:"race"`
		Sex              string         `json:"sex" db:"sex"`
		Age              int            `json:"ageInMonth" db:"age"`
		Description      string         `json:"description" db:"description"`
		Images           pq.StringArray `json:"imageUrls" db:"images"`
		IsAlreadyMatched bool           `json:"hasMatched" db:"is_already_matched"`
		CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
		UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
		DeletedAt        sql.NullTime   `json:"deletedAt,omitempty" db:"deleted_at"`
	}

	GetListCatRequest struct {
		ID          string `params:"id"`
		Limit       string `params:"limit" validate:"omitempty,number"`
		Offset      string `params:"offset" validate:"omitempty,number"`
		Race        string `params:"race" validate:"omitempty,validateRaces"`
		Sex         string `params:"sex" validate:"omitempty,oneof=male female"`
		Match       string `params:"isAlreadyMatched" validate:"omitempty,oneof=true false"`
		Age         string `params:"ageInMonth"`
		AgeValue    string
		AgeOperator string
		Owned       string `params:"owned" validate:"omitempty,oneof=true false"`
		Search      string `params:"search"`
		UserID      int
	}

	GetListCatResponse struct {
		ID               string         `json:"id" db:"id"`
		UserID           int            `json:"user_id" db:"user_id"`
		Name             string         `json:"name" db:"name"`
		Race             string         `json:"race" db:"race"`
		Sex              string         `json:"sex" db:"sex"`
		Age              int            `json:"ageInMonth" db:"age"`
		Description      string         `json:"description" db:"description"`
		Images           pq.StringArray `json:"imageUrls" db:"images"`
		IsAlreadyMatched bool           `json:"hasMatched" db:"is_already_matched"`
		CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
		UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
		DeletedAt        sql.NullTime   `json:"deletedAt,omitempty" db:"deleted_at"`
	}

	CreateCatRequest struct {
		UserID      int
		Name        string   `json:"name" validate:"required,min=1,max=30"`
		Race        string   `json:"race" validate:"required,validateRaces"`
		Sex         string   `json:"sex" validate:"required,oneof=male female"`
		AgeInMonth  int      `json:"ageInMonth" validate:"required,min=1,max=120082"`
		Description string   `json:"description" validate:"required,min=1,max=200"`
		ImageUrls   []string `json:"imageUrls" validate:"required,min=1,dive,min=1,url"`
	}

	CreateCatResponse struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
	}

	UpdateCatRequest struct {
		UserID      int
		ID          string   `param:"id" validate:"required"`
		Name        string   `json:"name" validate:"required,min=1,max=30"`
		Race        string   `json:"race" validate:"required,validateRaces"`
		Sex         string   `json:"sex" validate:"required,oneof=male female"`
		AgeInMonth  int      `json:"ageInMonth" validate:"required,min=1,max=120082"`
		Description string   `json:"description" validate:"required,min=1,max=200"`
		ImageUrls   []string `json:"imageUrls" validate:"required,min=1,dive,min=1,url"`
	}

	UpdateCatResponse struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Race        string    `json:"race"`
		Sex         string    `json:"sex"`
		AgeInMonth  int       `json:"ageInMonth"`
		Description string    `json:"description"`
		ImageUrls   []string  `json:"imageUrls"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	DeleteCatRequest struct {
		UserID int
		ID     string `param:"id" validate:"required"`
	}
)
