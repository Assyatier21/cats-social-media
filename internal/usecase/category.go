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

func (u *usecase) GetCategoryTree(ctx context.Context, req entity.GetCategoriesRequest) models.StandardResponseReq {
	var (
		categories = []entity.Category{}
	)

	categories, err := u.es.GetCategoryTree(ctx, req)
	if err != nil {
		log.Println("[Usecase][Category][GetCategoryTree] failed to get category tree, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_CATEGORIES, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.FAILED_GET_CATEGORIES, Data: categories, Error: nil}
}

func (u *usecase) GetCategoryDetails(ctx context.Context, req entity.GetCategoryDetailsRequest) models.StandardResponseReq {
	var (
		category = entity.Category{}
	)

	query := elastic.NewMatchQuery(constant.ID, req.ID)
	category, err := u.es.GetCategoryDetails(ctx, query)
	if err != nil {
		log.Println("[Usecase][Category][GetCategoryDetails] failed to get category details, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_CATEGORY_DETAILS, Error: err}
	}

	if category.ID == 0 {
		return models.StandardResponseReq{Code: http.StatusOK, Message: lib.ERR_DATA_NOT_FOUND, Error: nil}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.FAILED_GET_CATEGORY_DETAILS, Data: category, Error: nil}
}

func (u *usecase) InsertCategory(ctx context.Context, req entity.InsertCategoryRequest) models.StandardResponseReq {
	req.CreatedAt = constant.TimeNow
	req.UpdatedAt = constant.TimeNow

	category, err := u.repository.InsertCategory(ctx, req)
	if err != nil {
		log.Println("[Usecase][Category][InsertCategory] failed to insert category to postgres, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_CATEGORY_POSTGRES, Error: err}
	}

	err = u.es.InsertCategory(ctx, category)
	if err != nil {
		log.Println("[Usecase][Category][InsertCategory] failed to insert category to elastic, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_INSERT_CATEGORY_ELASTIC, Error: err}
	}

	category.CreatedAt = helper.FormattedTime(category.CreatedAt)
	category.UpdatedAt = helper.FormattedTime(category.UpdatedAt)

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_INSERT_CATEGORY, Data: category, Error: nil}
}

func (u *usecase) UpdateCategory(ctx context.Context, req entity.UpdateCategoryRequest) models.StandardResponseReq {
	req.UpdatedAt = constant.TimeNow

	err := u.repository.UpdateCategory(ctx, req)
	if err != nil {
		log.Println("[Usecase][Category][UpdateCategory] failed to update category to postgres, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_UPDATE_CATEGORY_POSTGRES, Error: err}
	}

	// Update Category To Elasticsearch
	err = u.es.UpdateCategory(ctx, req)
	if err != nil {
		log.Println("[Usecase][Category][UpdateCategory] failed to update category to elastic, err: ", err)
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_UPDATE_CATEGORY_ELASTIC, Error: err}
	}

	req.CreatedAt = helper.FormattedTime(req.CreatedAt)
	req.UpdatedAt = helper.FormattedTime(req.UpdatedAt)

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_CATEGORY, Data: req, Error: nil}
}

func (u *usecase) DeleteCategory(ctx context.Context, req entity.DeleteCategoryRequest) models.StandardResponseReq {
	var (
		responseChan = make(chan models.StandardResponseReq)
		wg           sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := u.repository.DeleteCategory(ctx, req)
		if err != nil {
			log.Println("[Usecase][Category][DeleteCategory] failed to delete category from postgres, err: ", err)
			responseChan <- models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_DELETE_CATEGORY_POSTGRES, Error: err}
		}
	}()

	go func() {
		defer wg.Done()
		err := u.es.DeleteCategory(ctx, req)
		if err != nil {
			log.Println("[Usecase][Category][DeleteCategory] failed to delete category from elastic, err: ", err)
			responseChan <- models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_DELETE_CATEGORY_ELASTIC, Error: err}
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
