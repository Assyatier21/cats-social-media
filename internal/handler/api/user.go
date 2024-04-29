package api

import (
	"net/http"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/helper"
	"github.com/backend-magang/cats-social-media/utils/pkg"
	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.CreateUserRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.RegisterUser(ctx, request)
	return helper.WriteResponse(c, resp)
}
