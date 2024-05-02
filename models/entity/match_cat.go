package entity

import (
	"database/sql"
	"time"
)

type MatchCatStatus string

const (
	MatchCatStatusApproved MatchCatStatus = "approved"
	MatchCatStatusPending  MatchCatStatus = "pending"
	MatchCatStatusRejected MatchCatStatus = "rejected"
)

type MatchCat struct {
	ID           int            `db:"id"`
	IssuedByID   int            `db:"issued_by_id"`
	TargetUserID int            `db:"target_user_id"`
	MatchCatID   int            `db:"match_cat_id"`
	UserCatID    int            `db:"user_cat_id"`
	Message      string         `db:"message"`
	Status       MatchCatStatus `db:"status"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	DeletedAt    sql.NullTime   `db:"deleted_at"`
}

type MatchCatRequest struct {
	UserID     int
	MatchCatID int    `json:"matchCatId" validate:"required"`
	UserCatID  int    `json:"userCatId" validate:"required"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}

type UpdateMatchCatRequest struct {
	UserID     int
	MatchCatID int `json:"matchId" validate:"required"`
}

type DeleteMatchCatRequest struct {
	UserID  int
	MatchID int `params:"id" validate:"required"`
}
type MatchApproveRequest struct {
	UserID  int
	MatchID int `json:"matchId" validate:"required"`
}
