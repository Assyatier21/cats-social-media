package router

import (
	"github.com/backend-magang/cats-social-media/internal/handler/api"
	_ "github.com/backend-magang/cats-social-media/middleware"
	"github.com/labstack/echo/v4"

	_ "github.com/backend-magang/cats-social-media/docs"
)

func InitRouter(server *echo.Echo, handler api.Handler) {
	InitCatRouter(server, handler)
}
