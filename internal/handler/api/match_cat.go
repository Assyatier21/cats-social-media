package api

import (
	"net/http"

	"github.com/backend-magang/cats-social-media/middleware"
	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/helper"
	"github.com/backend-magang/cats-social-media/utils/pkg"
	"github.com/labstack/echo/v4"
)

func (h *handler) MatchCat(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	user := middleware.ClaimToken(c)

	request := entity.MatchCatRequest{
		UserID: user.ID,
	}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.MatchCat(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) RejectMatchCat(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	user := middleware.ClaimToken(c)

	request := entity.UpdateMatchCatRequest{
		UserID: user.ID,
	}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.RejectMatchCat(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) GetListMatchCat(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	user := middleware.ClaimToken(c)

	request := entity.GetListMatchCatRequest{
		UserID: user.ID,
	}

	resp := h.usecase.GetListMatchCat(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) MatchApprove(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	user := middleware.ClaimToken(c)

	request := entity.MatchApproveRequest{
		UserID: user.ID,
	}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.MatchApprove(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) DeleteMatchCat(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	user := middleware.ClaimToken(c)

	request := entity.DeleteMatchCatRequest{
		UserID:  user.ID,
		MatchID: c.Param("id"),
	}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.DeleteMatchCat(ctx, request)
	return helper.WriteResponse(c, resp)
}
