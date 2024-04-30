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

	InsertUser(ctx context.Context, req entity.User) (err error)
	FindUserByEmail(ctx context.Context, email string) (result entity.User, err error)

	InsertMatchCat(ctx context.Context, req entity.MatchCat) (err error)
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
