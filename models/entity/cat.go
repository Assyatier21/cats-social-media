package entity

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Cat struct {
	ID               int            `json:"id" db:"id"`
	UserID           int            `json:"user_id" db:"user_id"`
	Name             string         `json:"name" db:"name"`
	Race             string         `json:"race" db:"race"`
	Sex              string         `json:"sex" db:"sex"`
	Age              int            `json:"age" db:"age"`
	Description      string         `json:"description" db:"description"`
	Images           pq.StringArray `json:"imageUrls" db:"images"`
	IsAlreadyMatched bool           `json:"isAlreadyMatched" db:"is_already_matched"`
	CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt        sql.NullTime   `json:"deletedAt,omitempty" db:"deleted_at"`
}

type GetListCatRequest struct {
	ID          string `params:"id"`
	Limit       int    `params:"limit"`
	Offset      int    `params:"offset"`
	Race        string `params:"race" validate:"omitempty,validateRaces"`
	Sex         string `params:"sex" validate:"omitempty,oneof=male female"`
	Match       string `params:"isAlreadyMatched"`
	Age         string `params:"ageInMonth" validate:"omitempty,validateAgeInMonth"`
	AgeValue    string
	AgeOperator string
	Owned       string `params:"owned"`
	Search      string `params:"search"`
	UserID      string
}
