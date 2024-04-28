package router

import (
	"github.com/backendmagang/project-1/internal/handler/api"
	"github.com/labstack/echo/v4"
)

func InitArticleRouter(e *echo.Echo, handler api.DeliveryHandler) {
	v2 := e.Group("/admin/v2")
	article := v2.Group("/article")

	article.GET("", handler.GetArticles)
	article.GET("/:id", handler.GetArticleDetails)
	article.POST("", handler.InsertArticle)
	article.PATCH("/:id", handler.UpdateArticle)
	article.DELETE("/:id", handler.DeleteArticle)
}
