package router

import (
	"github.com/backendmagang/project-1/internal/handler/api"
	"github.com/labstack/echo/v4"
)

func InitCategoryRouter(e *echo.Echo, handler api.DeliveryHandler) {
	v2 := e.Group("/admin/v2")
	category := v2.Group("/category")

	category.GET("", handler.GetCategories)
	category.GET("/:id", handler.GetCategoryDetails)
	category.POST("", handler.InsertCategory)
	category.PATCH("/:id", handler.UpdateCategory)
	category.DELETE("/:id", handler.DeleteCategory)
}
