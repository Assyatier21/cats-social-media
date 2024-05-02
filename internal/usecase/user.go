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
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) RegisterUser(ctx context.Context, req entity.CreateUserRequest) models.StandardResponseReq {
	var (
		newUser = entity.User{}
		now     = time.Now()
		usr     = entity.User{}
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

	usr, err = u.repository.InsertUser(ctx, newUser)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	token, _ := middleware.GenerateToken(entity.User{
		ID:    usr.ID,
		Email: usr.Email,
		Name:  usr.Name,
	})
	userJWT := entity.UserJWT{
		Email: newUser.Email,
		Name:  newUser.Name,
		Token: token,
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS_REGISTER_USER, Data: userJWT, Error: nil}
}

func (u *usecase) LoginUser(ctx context.Context, req entity.LoginUserRequest) models.StandardResponseReq {
	var (
		userJWT = entity.UserJWT{}
		user    = entity.User{}
		token   string
		err     error
	)

	user, err = u.repository.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED_LOGIN}
	}

	token, _ = middleware.GenerateToken(user)
	userJWT = entity.UserJWT{
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_LOGIN, Data: userJWT, Error: nil}
}
