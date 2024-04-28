package middleware

import (
	"github.com/backendmagang/project-1/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:  []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-App-Token, X-Client-Id"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	e.Use(middleware.CORS())
	e.Validator = &pkg.DataValidator{ValidatorData: pkg.SetupValidator()}
}
