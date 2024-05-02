package usecase

import (
	"context"
	"database/sql"
	"errors"
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
		Status:       entity.MatchCatStatusPending,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err = u.repository.InsertMatchCat(ctx, match)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS}
}

func (u *usecase) RejectMatchCat(ctx context.Context, req entity.UpdateMatchCatRequest) models.StandardResponseReq {
	var (
		now = time.Now()
	)

	matchCat, err := u.repository.FindMatchByID(ctx, req.MatchCatID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_MATCH_CAT_NOT_FOUND, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	if matchCat.Status == entity.MatchCatStatusRejected || matchCat.DeletedAt.Valid {
		return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED, Error: errors.New(constant.FAILED_MATCH_ID_INVALID)}
	}

	if matchCat.TargetUserID != req.UserID {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: errors.New(constant.FAILED_CAT_USER_UNAUTHORIZED)}
	}

	matchCat.Status = entity.MatchCatStatusRejected
	matchCat.UpdatedAt = now

	err = u.repository.UpdateMatchCat(ctx, matchCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS}
}

func (u *usecase) DeleteMatchCat(ctx context.Context, req entity.DeleteMatchCatRequest) models.StandardResponseReq {
	var (
		now = time.Now()
	)

	matchCat, err := u.repository.FindMatchByID(ctx, req.MatchID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_CAT_NOT_FOUND, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	if matchCat.Status == entity.MatchCatStatusApproved || matchCat.Status == entity.MatchCatStatusRejected || matchCat.DeletedAt.Valid {
		return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED, Error: errors.New(constant.FAILED_MATCH_ID_INVALID)}
	}

	if matchCat.IssuedByID != req.UserID {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: errors.New(constant.FAILED_CAT_USER_UNAUTHORIZED)}
	}

	matchCat.DeletedAt = sql.NullTime{Time: now, Valid: true}
	matchCat.UpdatedAt = now

	err = u.repository.UpdateMatchCat(ctx, matchCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS}
}
