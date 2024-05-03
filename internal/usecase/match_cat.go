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
	"github.com/backend-magang/cats-social-media/utils/pkg"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
)

func (u *usecase) MatchCat(ctx context.Context, req entity.MatchCatRequest) models.StandardResponseReq {
	var (
		g         errgroup.Group
		now       = time.Now()
		targetCat = entity.Cat{}
		userCat   = entity.Cat{}
		matchCats = []entity.MatchCat{}
		match     = entity.MatchCat{}
	)

	g.Go(func() (err1 error) {
		targetCat, err1 = u.repository.FindCatByID(ctx, cast.ToInt(req.MatchCatID))
		return err1
	})

	g.Go(func() (err2 error) {
		userCat, err2 = u.repository.FindCatByID(ctx, cast.ToInt(req.UserCatID))
		return err2
	})

	g.Go(func() (err3 error) {
		matchCats, err3 = u.repository.FindRequestedMatchCat(ctx, cast.ToInt(req.UserCatID), cast.ToInt(req.MatchCatID))
		return err3
	})

	if err := g.Wait(); err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_CAT_NOT_FOUND, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	if req.UserID != userCat.UserID {
		return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED, Error: errors.New(constant.FAILED_CAT_USER_UNAUTHORIZED)}
	}

	err := u.validateCatsWillBeMatched(ctx, req.UserID, targetCat, userCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: err}
	}

	if len(matchCats) > 0 {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: errors.New(constant.FAILED_REQUEST_MATCH_CATS)}
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

func (u *usecase) GetListMatchCat(ctx context.Context, req entity.GetListMatchCatRequest) models.StandardResponseReq {
	var (
		list_match_cat = []entity.GetListMatchCatQueryResponse{}
		resp           = []entity.GetListMatchCatResponse{}
		err            error
	)

	list_match_cat, err = u.repository.GetListMatchCat(ctx, req)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_MATCH_CATS, Error: err}
	}

	for _, match_cat := range list_match_cat {
		data := buildResponseListMatchCat(match_cat)
		resp = append(resp, data)
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: resp}
}

func (u *usecase) MatchApprove(ctx context.Context, req entity.MatchApproveRequest) models.StandardResponseReq {
	matchCat, err := u.repository.FindMatchCatByID(ctx, req.MatchID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_GET_MATCH_ID, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	err = u.validateMatchCat(ctx, req.UserID, matchCat)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: err}
	}

	if err = pkg.WithTransaction(ctx, u.cfg.SqlTrx, func(ctx context.Context) (err error) {
		err = u.repository.UpdateCatsMatch(ctx, matchCat)
		if err != nil {
			return
		}

		err = u.repository.ApproveMatch(ctx, matchCat.ID)
		if err != nil {
			return
		}

		err = u.repository.DeleteOtherMatch(ctx, matchCat.UserCatID, matchCat.ID)
		if err != nil {
			return
		}

		err = u.repository.DeleteOtherMatch(ctx, matchCat.MatchCatID, matchCat.ID)
		if err != nil {
			return
		}

		return nil
	}); err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Data: "LALA"}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_APPROVE_MATCH, Data: nil}
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
