package postgres

import (
	"context"
	"log"

	"github.com/backendmagang/project-1/models/entity"
)

func (r *repository) buildInsertedArticleResponse(req entity.InsertArticleRequest) (article entity.ArticleResponse) {
	return entity.ArticleResponse{
		ID:           req.ID,
		Title:        req.Title,
		Slug:         req.Slug,
		HTMLContent:  req.HTMLContent,
		Metadata:     req.Metadata,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
		CategoryList: r.buildCategoryList(req.CategoryIDs),
	}
}

func (r *repository) buildUpdatedArticleResponse(req entity.UpdateArticleRequest) (article entity.ArticleResponse) {
	return entity.ArticleResponse{
		ID:           req.ID,
		Title:        req.Title,
		Slug:         req.Slug,
		HTMLContent:  req.HTMLContent,
		Metadata:     req.Metadata,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
		CategoryList: r.buildCategoryList(req.CategoryIDs),
	}
}

func (r *repository) buildCategoryList(categoryIDs []int) []entity.CategoryResponse {
	categories, err := r.GetCategoriesByIDs(context.Background(), categoryIDs)
	if err != nil {
		log.Println("[Repository][Postgres][buildCategoryList] error failed to get category by ids, err: ", err)
		return nil
	}

	return categories
}
