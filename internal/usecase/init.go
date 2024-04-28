package usecase

import (
	"context"

	"github.com/backend-magang/cats-social-media/internal/repository/postgres"
	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
)

type UsecaseHandler interface {
	GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq
}

type usecase struct {
	repository postgres.RepositoryHandler
}

func NewUsecase(repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		repository: repository,
	}
}
