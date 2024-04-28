package postgres

import (
	"context"
	"encoding/json"
	"log"

	"github.com/backendmagang/project-1/internal/repository/postgres/queries"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/backendmagang/project-1/models/lib"
	"github.com/lib/pq"
)

func (r *repository) GetArticles(ctx context.Context, req entity.GetArticlesRequest) ([]entity.ArticleResponse, error) {
	var (
		articles = []entity.ArticleResponse{}
	)

	rows, err := r.db.Query(queries.GET_ARTICLES, req.Limit, req.Offset)
	if err != nil {
		log.Println("[Repository][Postgres][GetArticles] error failed to exec query, err: ", err)
		return articles, nil

	}
	defer rows.Close()

	for rows.Next() {
		var article entity.ArticleResponse
		var categoryJSON string

		err := rows.Scan(&article.ID, &article.Title, &article.Slug, &article.HTMLContent, &article.Metadata, &article.CreatedAt, &article.UpdatedAt, &categoryJSON)
		if err != nil {
			log.Println("[Repository][Postgres][GetArticles] error failed to scan data, err: ", err)
			return articles, nil
		}

		err = json.Unmarshal([]byte(categoryJSON), &article.CategoryList)
		if err != nil {
			log.Println("[Repository][Postgres][GetArticles] error failed to unmarshal categories, err: ", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *repository) GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) (entity.ArticleResponse, error) {
	var (
		article      = entity.ArticleResponse{}
		categoryJSON string
		categories   []entity.CategoryResponse
	)

	err := r.db.QueryRow(queries.GET_ARTICLE_DETAILS, req.ID).Scan(&article.ID, &article.Title, &article.Slug, &article.HTMLContent, &article.Metadata, &article.CreatedAt, &article.UpdatedAt, &categoryJSON)
	if err != nil {
		log.Println("[Repository][Postgres][GetArticleDetails] error failed to scan result, err: ", err)
		return article, err
	}

	if err := json.Unmarshal([]byte(categoryJSON), &categories); err != nil {
		log.Println("[Repository][Postgres][GetArticleDetails] error failed to unmarshal categories, err: ", err)
		return article, err
	}

	article.CategoryList = categories
	return article, err
}

func (r *repository) InsertArticle(ctx context.Context, article entity.InsertArticleRequest) (entity.ArticleResponse, error) {
	_, err := r.db.Exec(queries.INSERT_ARTICLE, article.ID, article.Title, article.Slug, article.HTMLContent, pq.Array(article.CategoryIDs), article.Metadata, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][Postgres][InsertArticle] error failed to insert article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	articleResponse := r.buildInsertedArticleResponse(article)
	return articleResponse, nil
}

func (r *repository) UpdateArticle(ctx context.Context, article entity.UpdateArticleRequest) (entity.ArticleResponse, error) {
	result, err := r.db.Exec(queries.UPDATE_ARTICLE, article.Title, article.Slug, article.HTMLContent, pq.Array(article.CategoryIDs), article.Metadata, article.CreatedAt, article.UpdatedAt, article.ID)
	if err != nil {
		log.Println("[Repository][Postgres][UpdateArticle] error failed to update article, err: ", err)
		return entity.ArticleResponse{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.ArticleResponse{}, lib.ErrorNoRowsAffected
	}

	articleResponse := r.buildUpdatedArticleResponse(article)

	return articleResponse, nil
}

func (r *repository) DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) error {
	result, err := r.db.Exec(queries.DELETE_ARTICLE, req.ID)
	if err != nil {
		log.Println("[Repository][Postgres][DeleteArticle] error failed to delete article, err: ", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return lib.ErrorNoRowsAffected
	}

	return nil
}
