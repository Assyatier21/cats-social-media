package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"golang.org/x/sync/errgroup"
)

func (u *usecase) MatchCat(ctx context.Context, req entity.MatchCatRequest) models.StandardResponseReq {
	var (
		g         errgroup.Group
		now       = time.Now()
		targetCat = entity.Cat{}
		userCat   = entity.Cat{}
		match     = entity.MatchCat{}
	)

	g.Go(func() (err1 error) {
		targetCat, err1 = u.repository.FindCatByID(ctx, req.MatchCatID)
		return err1
	})

	g.Go(func() (err2 error) {
		userCat, err2 = u.repository.FindCatByID(ctx, req.UserCatID)
		return err2
	})

	if err := g.Wait(); err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_CAT_NOT_FOUND, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	err := u.validateMatchCat(ctx, req.UserID, targetCat, userCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: err}
	}

	match = entity.MatchCat{
		IssuedByID:   req.UserID,
		TargetUserID: targetCat.UserID,
		MatchCatID:   targetCat.ID,
		UserCatID:    userCat.ID,
		Message:      req.Message,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err = u.repository.InsertMatchCat(ctx, match)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS}
}

func (u *usecase) GetListMatchCat(ctx context.Context) models.StandardResponseReq {
	result, err := u.repository.GetListMatchCat(ctx)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_CATS, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: result}
}
