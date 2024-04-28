package api

import (
	"github.com/backendmagang/project-1/internal/usecase"
	"github.com/labstack/echo/v4"
)

type DeliveryHandler interface {
	GetArticles(c echo.Context) (err error)
	GetArticleDetails(c echo.Context) (err error)
	InsertArticle(c echo.Context) (err error)
	UpdateArticle(c echo.Context) (err error)
	DeleteArticle(c echo.Context) (err error)

	GetCategories(c echo.Context) (err error)
	GetCategoryDetails(c echo.Context) (err error)
	InsertCategory(c echo.Context) (err error)
	UpdateCategory(c echo.Context) (err error)
	DeleteCategory(c echo.Context) (err error)
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func NewHandler(usecase usecase.UsecaseHandler) DeliveryHandler {
	return &handler{
		usecase: usecase,
	}
}
