package router

import (
	"github.com/backendmagang/project-1/internal/handler/api"
	_ "github.com/backendmagang/project-1/middleware"
	"github.com/labstack/echo/v4"

	_ "github.com/backendmagang/project-1/docs"
)

func InitRouter(server *echo.Echo, handler api.DeliveryHandler) {
	InitArticleRouter(server, handler)
	InitCategoryRouter(server, handler)
}
