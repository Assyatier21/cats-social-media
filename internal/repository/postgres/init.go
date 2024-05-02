package postgres

import (
	"context"

	"github.com/backend-magang/cats-social-media/config"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryHandler interface {
	GetListCat(ctx context.Context, req entity.GetListCatRequest) (result []entity.Cat, err error)
	InsertCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	UpdateCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	FindUserCatByID(ctx context.Context, userId int, catId int) (result entity.Cat, err error)
	FindCatByID(ctx context.Context, id int) (entity.Cat, error)

	InsertUser(ctx context.Context, req entity.User) (result entity.User, err error)
	FindUserByEmail(ctx context.Context, email string) (result entity.User, err error)

	FindRequestedMatch(ctx context.Context, catId int) (result []entity.MatchCat, err error)
	InsertMatchCat(ctx context.Context, req entity.MatchCat) (err error)
	FindMatchCatByID(ctx context.Context, id int) (result entity.MatchCat, err error)
	ApproveMatch(ctx context.Context, matchCatId int) (err error)
	UpdateCatsMatch(ctx context.Context, matchCat entity.MatchCat) (err error)
	DeleteOtherMatch(ctx context.Context, catId int, matchCatId int) (err error)
}

type repository struct {
	cfg    config.Config
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewRepository(cfg config.Config, db *sqlx.DB, log *logrus.Logger) RepositoryHandler {
	return &repository{
		cfg:    cfg,
		db:     db,
		logger: log,
	}
}
