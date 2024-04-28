package router

import (
	"github.com/backend-magang/cats-social-media/internal/handler/api"
	"github.com/labstack/echo/v4"
)

func InitCatRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	cat := v1.Group("/cat")

	_ = cat

	cat.GET("", handler.GetListCat)
	// article.GET("/:id", handler.GetArticleDetails)
	// article.POST("", handler.InsertArticle)
	// article.PATCH("/:id", handler.UpdateArticle)
	// article.DELETE("/:id", handler.DeleteArticle)
}
