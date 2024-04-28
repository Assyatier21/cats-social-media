package usecase

import (
	"context"

	elastic "github.com/backendmagang/project-1/internal/repository/elasticsearch"
	"github.com/backendmagang/project-1/internal/repository/postgres"
	"github.com/backendmagang/project-1/models"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/redis/go-redis/v9"
)

type UsecaseHandler interface {
	GetArticles(ctx context.Context, req entity.GetArticlesRequest) models.StandardResponseReq
	GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) models.StandardResponseReq
	InsertArticle(ctx context.Context, req entity.InsertArticleRequest) models.StandardResponseReq
	UpdateArticle(ctx context.Context, req entity.UpdateArticleRequest) models.StandardResponseReq
	DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) models.StandardResponseReq

	GetCategoryTree(ctx context.Context, req entity.GetCategoriesRequest) models.StandardResponseReq
	GetCategoryDetails(ctx context.Context, req entity.GetCategoryDetailsRequest) models.StandardResponseReq
	InsertCategory(ctx context.Context, req entity.InsertCategoryRequest) models.StandardResponseReq
	UpdateCategory(ctx context.Context, req entity.UpdateCategoryRequest) models.StandardResponseReq
	DeleteCategory(ctx context.Context, req entity.DeleteCategoryRequest) models.StandardResponseReq
}

type usecase struct {
	repository postgres.RepositoryHandler
	es         elastic.ElasticHandler
	redis      *redis.Client
}

func NewUsecase(repository postgres.RepositoryHandler, es elastic.ElasticHandler, redis *redis.Client) UsecaseHandler {
	return &usecase{
		repository: repository,
		es:         es,
		redis:      redis,
	}
}
