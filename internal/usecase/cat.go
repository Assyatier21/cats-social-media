package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"github.com/spf13/cast"
)

func (u *usecase) GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq {
	if err := builFilterAgeRequest(&req); err != nil {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: err}
	}

	cats, err := u.repository.GetListCat(ctx, req)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_CATS, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: cats}
}

func (u *usecase) CreateCat(ctx context.Context, req entity.CreateCatRequest) models.StandardResponseReq {
	now := time.Now()

	newCat := entity.Cat{
		UserID:           req.UserID,
		Name:             req.Name,
		Race:             req.Race,
		Sex:              req.Sex,
		Age:              req.AgeInMonth,
		Description:      req.Description,
		Images:           req.ImageUrls,
		IsAlreadyMatched: false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	cat, err := u.repository.InsertCat(ctx, newCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	resp := entity.CreateCatResponse{
		ID:        cast.ToString(cat.ID),
		CreatedAt: cat.CreatedAt,
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS_ADD_CAT, Data: resp, Error: nil}
}

func (u *usecase) UpdateCat(ctx context.Context, req entity.UpdateCatRequest) models.StandardResponseReq {
	var (
		now    = time.Now()
		userId = req.UserID
		catId  = cast.ToInt(req.ID)
	)

	// find user's cat by id
	cat, err := u.repository.FindUserCatByID(ctx, userId, catId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_GET_USER_CAT, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// if sex changed
	if req.Sex != cat.Sex {
		matchCats, err := u.repository.FindRequestedMatch(ctx, catId)
		if err != nil && err != sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
		}

		// if there is a match request
		if len(matchCats) > 0 {
			return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.HAS_REQUESTED_MATCH, Error: err}
		}
	}

	updatedCat := entity.Cat{
		UserID:      userId,
		ID:          catId,
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		Age:         req.AgeInMonth,
		Description: req.Description,
		Images:      req.ImageUrls,
		UpdatedAt:   now,
	}

	cat, err = u.repository.UpdateCat(ctx, updatedCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	resp := entity.UpdateCatResponse{
		ID:          cast.ToString(cat.ID),
		Name:        cat.Name,
		Race:        cat.Race,
		Sex:         cat.Sex,
		AgeInMonth:  cat.Age,
		Description: cat.Description,
		ImageUrls:   cat.Images,
		CreatedAt:   cat.CreatedAt,
		UpdatedAt:   cat.UpdatedAt,
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_CAT, Data: resp}
}
