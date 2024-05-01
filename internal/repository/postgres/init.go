package postgres

import (
	"context"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryHandler interface {
	GetListCat(ctx context.Context, req entity.GetListCatRequest) (result []entity.Cat, err error)
	FindCatByID(ctx context.Context, id int) (entity.Cat, error)
	FindUserCatByID(ctx context.Context, userId int, catId int) (result entity.Cat, err error)
	InsertCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	UpdateCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)

	FindUserByEmail(ctx context.Context, email string) (result entity.User, err error)
	InsertUser(ctx context.Context, req entity.User) (result entity.User, err error)

	FindRequestedMatch(ctx context.Context, catId int) (result []entity.MatchCat, err error)
	FindMatchByID(ctx context.Context, id int) (entity.MatchCat, error)
	InsertMatchCat(ctx context.Context, req entity.MatchCat) (err error)
	UpdateMatchCat(ctx context.Context, req entity.MatchCat) (err error)
}

type repository struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewRepository(db *sqlx.DB, log *logrus.Logger) RepositoryHandler {
	return &repository{
		db:     db,
		logger: log,
	}
}
