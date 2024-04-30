package router

import (
	"github.com/backend-magang/cats-social-media/internal/handler/api"
	"github.com/backend-magang/cats-social-media/middleware"
	"github.com/labstack/echo/v4"
)

func InitCatRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	cat := v1.Group("/cat", middleware.TokenValidationMiddleware())

	cat.GET("", handler.GetListCat)
	cat.POST("", handler.CreateCat)
	cat.PUT("/:id", handler.UpdateCat)
	// article.GET("/:id", handler.GetArticleDetails)
	// article.POST("", handler.InsertArticle)
	// article.PATCH("/:id", handler.UpdateArticle)
	// article.DELETE("/:id", handler.DeleteArticle)
}
