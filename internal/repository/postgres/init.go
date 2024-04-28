package postgres

import (
	"context"
	"database/sql"

	"github.com/backendmagang/project-1/models/entity"
)

type RepositoryHandler interface {
	GetArticles(ctx context.Context, req entity.GetArticlesRequest) ([]entity.ArticleResponse, error)
	GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) (entity.ArticleResponse, error)
	InsertArticle(ctx context.Context, article entity.InsertArticleRequest) (entity.ArticleResponse, error)
	UpdateArticle(ctx context.Context, article entity.UpdateArticleRequest) (entity.ArticleResponse, error)
	DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) error

	GetCategoryTree(ctx context.Context, req entity.GetCategoriesRequest) ([]entity.Category, error)
	GetCategoryDetails(ctx context.Context, req entity.GetCategoryDetailsRequest) (entity.Category, error)
	InsertCategory(ctx context.Context, category entity.InsertCategoryRequest) (entity.InsertCategoryRequest, error)
	UpdateCategory(ctx context.Context, category entity.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, req entity.DeleteCategoryRequest) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryHandler {
	return &repository{
		db: db,
	}
}
