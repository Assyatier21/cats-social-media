package usecase

import (
	"context"

	"github.com/backend-magang/cats-social-media/config"
	"github.com/backend-magang/cats-social-media/internal/repository/postgres"
	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/sirupsen/logrus"
)

type UsecaseHandler interface {
	GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq
	CreateCat(ctx context.Context, req entity.CreateCatRequest) models.StandardResponseReq
	UpdateCat(ctx context.Context, req entity.UpdateCatRequest) models.StandardResponseReq

	MatchCat(ctx context.Context, req entity.MatchCatRequest) models.StandardResponseReq
	RejectMatchCat(ctx context.Context, req entity.UpdateMatchCatRequest) models.StandardResponseReq
	DeleteMatchCat(ctx context.Context, req entity.DeleteMatchCatRequest) models.StandardResponseReq
	GetListMatchCat(ctx context.Context, req entity.GetListMatchCatRequest) models.StandardResponseReq
	MatchApprove(ctx context.Context, req entity.MatchApproveRequest) models.StandardResponseReq

	RegisterUser(ctx context.Context, req entity.CreateUserRequest) models.StandardResponseReq
	LoginUser(ctx context.Context, req entity.LoginUserRequest) models.StandardResponseReq
}

type usecase struct {
	cfg        config.Config
	logger     *logrus.Logger
	repository postgres.RepositoryHandler
}

func NewUsecase(cfg config.Config, log *logrus.Logger, repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		cfg:        cfg,
		logger:     log,
		repository: repository,
	}
}
