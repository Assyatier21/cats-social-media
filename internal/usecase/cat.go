package usecase

import (
	"context"
	"net/http"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
)

func (u *usecase) GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq {
	var (
		cats = []entity.Cat{}
	)

	if req.Age != "" {
		builFilterAgeRequest(&req)
	}

	cats, err := u.repository.GetListCat(ctx, req)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_CATS, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: cats}
}
