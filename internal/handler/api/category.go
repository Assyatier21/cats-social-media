package api

import (
	"net/http"

	"github.com/backendmagang/project-1/models"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/backendmagang/project-1/utils/helper"
	"github.com/backendmagang/project-1/utils/pkg"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetCategories(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.GetCategoriesRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.GetCategoryTree(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) GetCategoryDetails(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.GetCategoryDetailsRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.GetCategoryDetails(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) InsertCategory(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.InsertCategoryRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.InsertCategory(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) UpdateCategory(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.UpdateCategoryRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.UpdateCategory(ctx, req)
	return helper.WriteResponse(c, resp)
}

func (h *handler) DeleteCategory(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	req := entity.DeleteCategoryRequest{}
	err = pkg.BindValidate(c, &req)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.DeleteCategory(ctx, req)
	return helper.WriteResponse(c, resp)
}
