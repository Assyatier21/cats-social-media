package usecase

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/backendmagang/project-1/models"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/backendmagang/project-1/models/lib"
	"github.com/backendmagang/project-1/utils/constant"
	"github.com/backendmagang/project-1/utils/helper"
	"github.com/olivere/elastic/v7"
)

func (u *usecase) GetArticles(ctx context.Context, req entity.GetArticlesRequest) models.StandardResponseReq {
	var (
		articles = []entity.ArticleResponse{}
	)

	req = helper.GeArticleRequestValidation(req)
	articles, err := u.es.GetArticles(ctx, req)
	if err != nil {
		log.Println("[Usecase][Article][Article][GetArticles] failed to get list of articles, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_ARTICLES, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_GET_ARTICLES, Data: articles, Error: nil}
}

func (u *usecase) GetArticleDetails(ctx context.Context, req entity.GetArticleDetailsRequest) models.StandardResponseReq {
	var (
		article = entity.ArticleResponse{}
	)

	query := elastic.NewMatchQuery(constant.ID, req.ID)
	article, err := u.es.GetArticleDetails(ctx, query)
	if err != nil {
		log.Println("[Usecase][Article][Article][GetArticleDetails] failed to get article details, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_ARTICLE_DETAILS, Error: err}
	}

	if article.ID == "" {
		return models.StandardResponseReq{Code: http.StatusOK, Message: lib.ERR_DATA_NOT_FOUND, Error: nil}
	}

	article = helper.FormatTimeArticleResponse(article)
	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.FAILED_GET_ARTICLE_DETAILS, Data: article, Error: nil}
}

func (u *usecase) InsertArticle(ctx context.Context, req entity.InsertArticleRequest) models.StandardResponseReq {
	// Insert UUID Article
	req.ID = helper.GenerateUUIDString()
	req.CreatedAt = constant.TimeNow
	req.UpdatedAt = constant.TimeNow

	articleResponse, err := u.repository.InsertArticle(ctx, req)
	if err != nil {
		log.Println("[Usecase][Article][InsertArticle] failed to insert article to postgres, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE_POSTGRES, Error: err}
	}

	err = u.es.InsertArticle(ctx, articleResponse)
	if err != nil {
		log.Println("[Usecase][Article][InsertArticle] failed to insert article to elastic, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_ARTICLE_ELASTIC, Error: err}
	}

	helper.FormatTimeArticleResponse(articleResponse)
	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_INSERT_ARTICLE, Data: articleResponse, Error: nil}
}

func (u *usecase) UpdateArticle(ctx context.Context, req entity.UpdateArticleRequest) models.StandardResponseReq {
	var (
		article = entity.ArticleResponse{}
		err     error
	)

	// Change updated_at time
	req.UpdatedAt = constant.TimeNow

	article, err = u.repository.UpdateArticle(ctx, req)
	if err != nil {
		log.Println("[Usecase][Article][UpdateArticle] failed to update article to postgres, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_UPDATE_ARTICLE_POSTGRES, Error: err}
	}

	// Update Article To Elasticsearch
	err = u.es.UpdateArticle(ctx, article)
	if err != nil {
		log.Println("[Usecase][Article][UpdateArticle] failed to update article to elastic, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_UPDATE_ARTICLE_ELASTIC, Error: err}
	}

	article.CreatedAt = helper.FormattedTime(article.CreatedAt)
	article.UpdatedAt = helper.FormattedTime(constant.TimeNow)

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_ARTICLE, Data: article, Error: nil}
}

func (u *usecase) DeleteArticle(ctx context.Context, req entity.DeleteArticleRequest) models.StandardResponseReq {
	var (
		responseChan = make(chan models.StandardResponseReq)
		wg           sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := u.repository.DeleteArticle(ctx, req)
		if err != nil {
			log.Println("[Usecase][Article][DeleteArticle] failed to delete article from postgres, err: ", err)
			responseChan <- models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_DELETE_ARTICLE_POSTGRES, Error: err}
		}
	}()

	go func() {
		defer wg.Done()
		err := u.es.DeleteArticle(ctx, req)
		if err != nil {
			log.Println("[Usecase][Article][DeleteArticle] failed to delete article from elastic, err: ", err)
			responseChan <- models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_DELETE_ARTICLE_ELASTIC, Error: err}
		}
	}()

	go func() {
		wg.Wait()
		close(responseChan)
	}()

	for response := range responseChan {
		if response.Error != nil {
			return response
		}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_DELETE_ARTICLE, Error: nil}
}
