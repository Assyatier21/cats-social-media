package postgres

import (
	"context"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/jmoiron/sqlx"
)

type RepositoryHandler interface {
	GetListCat(ctx context.Context, req entity.GetListCatRequest) (result []entity.Cat, err error)
	InsertCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	UpdateCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	FindUserCatByID(ctx context.Context, userId int, catId int) (result entity.Cat, err error)

	InsertUser(ctx context.Context, req entity.User) (err error)
	FindUserByEmail(ctx context.Context, email string) (result entity.User, err error)

	FindRequestedMatch(ctx context.Context, catId int) (result []entity.MatchCat, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) RepositoryHandler {
	return &repository{
		db: db,
	}
}
