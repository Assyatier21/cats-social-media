package api

import (
	"net/http"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/helper"
	"github.com/backend-magang/cats-social-media/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (h *handler) GetListCat(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.GetListCatRequest{
		ID:     c.QueryParam("id"),
		Limit:  cast.ToInt(c.QueryParam("limit")),
		Offset: cast.ToInt(c.QueryParam("offset")),
		Race:   c.QueryParam("race"),
		Sex:    c.QueryParam("sex"),
		Match:  c.QueryParam("isAlreadyMatched"),
		Age:    c.QueryParam("ageInMonth"),
		Owned:  c.QueryParam("owned"),
		Search: c.QueryParam("search"),
	}

	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.GetListCat(ctx, request)
	return helper.WriteResponse(c, resp)
}
