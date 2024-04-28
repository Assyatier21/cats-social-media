package helper

import (
	"net/http"

	"github.com/backendmagang/project-1/models"
	"github.com/backendmagang/project-1/utils/constant"

	"github.com/labstack/echo/v4"
)

func WriteResponse(c echo.Context, req models.StandardResponseReq) error {
	var status = constant.SUCCESS

	if req.Code > 299 {
		status = constant.FAILED
	}
	var errResp interface{}
	if req.Error != nil {
		errResp = req.Error.Error()
	}

	if req.Message == "" {
		req.Message = http.StatusText(req.Code)
	}

	return c.JSON(req.Code, models.StandardResponse{
		Code:    req.Code,
		Status:  status,
		Message: req.Message,
		Data:    req.Data,
		Error:   errResp,
	})
}
