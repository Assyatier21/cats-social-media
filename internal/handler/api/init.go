package api

import (
	"github.com/backend-magang/cats-social-media/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetListCat(c echo.Context) (err error)

	RegisterUser(c echo.Context) (err error)
	LoginUser(c echo.Context) (err error)
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func NewHandler(usecase usecase.UsecaseHandler) Handler {
	return &handler{
		usecase: usecase,
	}
}
