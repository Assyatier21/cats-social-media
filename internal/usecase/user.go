package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/backend-magang/cats-social-media/middleware"
	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"github.com/backend-magang/cats-social-media/utils/helper"
	"github.com/spf13/cast"
)

func (u *usecase) RegisterUser(ctx context.Context, req entity.CreateUserRequest) models.StandardResponseReq {
	var (
		newUser = entity.User{}
		now     = time.Now()
	)

	// Check If Phone Already Registered
	user, err := u.repository.FindUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	if user.Email != "" {
		return models.StandardResponseReq{Code: http.StatusConflict, Message: constant.EMAIL_REGISTERED, Error: nil}
	}

	newUser = entity.User{
		Email:     req.Email,
		Name:      req.Name,
		Password:  helper.HashPassword(req.Password, cast.ToInt(u.cfg.BCryptSalt)),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = u.repository.InsertUser(ctx, newUser)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	token, _ := middleware.GenerateToken(newUser)
	userJWT := entity.UserJWT{
		Email: newUser.Email,
		Name:  newUser.Name,
		Token: token,
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS_REGISTER_USER, Data: userJWT, Error: nil}
}
