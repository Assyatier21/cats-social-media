package entity

import (
	"database/sql"
	"time"
)

type (
	MatchCat struct {
		ID           int          `db:"id"`
		IssuedByID   int          `db:"issued_by_id"`
		TargetUserID int          `db:"target_user_id"`
		MatchCatID   int          `db:"match_cat_id"`
		UserCatID    int          `db:"user_cat_id"`
		Message      string       `db:"message"`
		CreatedAt    time.Time    `db:"created_at"`
		UpdatedAt    time.Time    `db:"updated_at"`
		DeletedAt    sql.NullTime `db:"deleted_at"`
	}

	MatchCatRequest struct {
		UserID     int
		MatchCatID int    `json:"matchCatId" validate:"required"`
		UserCatID  int    `json:"userCatId" validate:"required"`
		Message    string `json:"message" validate:"required,min=5,max=120"`
	}
)
