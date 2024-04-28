package api

import (
	"net/http"

	"github.com/backendmagang/project-1/models"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/backendmagang/project-1/utils/helper"
	"github.com/backendmagang/project-1/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (h *handler) GetArticles(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.GetArticlesRequest{
		Limit:   cast.ToInt(c.QueryParam("limit")),
		Offset:  cast.ToInt(c.QueryParam("offset")),
		SortBy:  c.QueryParam("sort_by"),
		OrderBy: c.QueryParam("order_by"),
	}

	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.GetArticles(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) GetArticleDetails(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.GetArticleDetailsRequest{ID: c.Param("id")}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.GetArticleDetails(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) InsertArticle(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.InsertArticleRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.InsertArticle(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) UpdateArticle(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.UpdateArticleRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.UpdateArticle(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) DeleteArticle(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.DeleteArticleRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.DeleteArticle(ctx, req)
	return helper.WriteResponse(c, resp)
}
