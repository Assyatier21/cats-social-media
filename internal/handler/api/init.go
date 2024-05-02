package api

import (
	"github.com/backend-magang/cats-social-media/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	GetListCat(c echo.Context) (err error)
	CreateCat(c echo.Context) (err error)
	UpdateCat(c echo.Context) (err error)

	MatchCat(c echo.Context) (err error)
	RejectMatchCat(c echo.Context) (err error)
	DeleteMatchCat(c echo.Context) (err error)
	MatchApprove(c echo.Context) (err error)

	RegisterUser(c echo.Context) (err error)
	LoginUser(c echo.Context) (err error)
}

type handler struct {
	logger  *logrus.Logger
	usecase usecase.UsecaseHandler
}

func NewHandler(log *logrus.Logger, usecase usecase.UsecaseHandler) Handler {
	return &handler{
		logger:  log,
		usecase: usecase,
	}
}
