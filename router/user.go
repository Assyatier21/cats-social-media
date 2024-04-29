package router

import (
	"github.com/backend-magang/cats-social-media/internal/handler/api"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	user := v1.Group("/user")

	user.POST("/register", handler.RegisterUser)
}
