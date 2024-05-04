package entity

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
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

type MatchCatWithUserAndCats struct {
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

	IssuedBy       User `db:"-"`
	MatchCatDetail Cat  `db:"match_cat"`
	UserCatDetail  Cat  `db:"user_cat"`
}

type MatchCatRequest struct {
	UserID     int
	MatchCatID string `json:"matchCatId" validate:"required"`
	UserCatID  string `json:"userCatId" validate:"required"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}

type UpdateMatchCatRequest struct {
	UserID     int
	MatchCatID string `json:"matchId" validate:"required"`
}

type DeleteMatchCatRequest struct {
	UserID  int
	MatchID int `params:"id" validate:"required"`
}
type GetListMatchCatRequest struct {
	UserID int
}

type GetListMatchCatResponse struct {
	ID             string       `json:"id"`
	IssuedBy       IssuedByData `json:"issuedBy"`
	MatchCatDetail CatDetail    `json:"matchCatDetail"`
	UserCatDetail  CatDetail    `json:"userCatDetail"`
	Message        string       `json:"message"`
	CreatedAt      time.Time    `json:"createdAt"`
}

type IssuedByData struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CatDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GetListMatchCatQueryResponse struct {
	ID                int       `db:"match_id" json:"match_id"`
	Message           string    `db:"message" json:"message"`
	CreatedAt         time.Time `db:"match_created_at" json:"match_created_at"`
	IssuedByName      string    `db:"issuedby_name" json:"issuedby_name"`
	IssuedByEmail     string    `db:"issuedby_email" json:"issuedby_email"`
	IssuedByCreatedAt time.Time `db:"issuedby_createdat" json:"issuedby_createdat"`

	MatchCatID          int            `db:"match_cat_id" json:"match_cat_id"`
	MatchCatName        string         `db:"match_cat_name" json:"match_cat_name"`
	MatchCatRace        string         `db:"match_cat_race" json:"match_cat_race"`
	MatchCatSex         string         `db:"match_cat_sex" json:"match_cat_sex"`
	MatchCatDescription string         `db:"match_cat_description" json:"match_cat_description"`
	MatchCatAge         int            `db:"match_cat_age" json:"match_cat_age"`
	MatchCatImages      pq.StringArray `db:"match_cat_images" json:"match_cat_images"`
	MatchCatHasMatched  bool           `db:"match_cat_hasmatched" json:"match_cat_hasmatched"`
	MatchCatCreatedAt   time.Time      `db:"match_cat_createdat" json:"match_cat_createdat"`

	UserCatID          int            `db:"user_cat_id" json:"user_cat_id"`
	UserCatName        string         `db:"user_cat_name" json:"user_cat_name"`
	UserCatRace        string         `db:"user_cat_race" json:"user_cat_race"`
	UserCatSex         string         `db:"user_cat_sex" json:"user_cat_sex"`
	UserCatDescription string         `db:"user_cat_description" json:"user_cat_description"`
	UserCatAge         int            `db:"user_cat_age" json:"user_cat_age"`
	UserCatImages      pq.StringArray `db:"user_cat_images" json:"user_cat_images"`
	UserCatHasMatched  bool           `db:"user_cat_hasmatched" json:"user_cat_hasmatched"`
	UserCatCreatedAt   time.Time      `db:"user_cat_createdat" json:"user_cat_createdat"`
}
type MatchApproveRequest struct {
	UserID  int
	MatchID string `json:"matchId" validate:"required"`
}
