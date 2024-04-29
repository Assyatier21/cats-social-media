package usecase

import (
	"context"

	"github.com/backend-magang/cats-social-media/config"
	"github.com/backend-magang/cats-social-media/internal/repository/postgres"
	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
)

type UsecaseHandler interface {
	GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq

	RegisterUser(ctx context.Context, req entity.CreateUserRequest) models.StandardResponseReq
	LoginUser(ctx context.Context, req entity.LoginUserRequest) models.StandardResponseReq
}

type usecase struct {
	cfg        config.Config
	repository postgres.RepositoryHandler
}

func NewUsecase(cfg config.Config, repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		cfg:        cfg,
		repository: repository,
	}
}
