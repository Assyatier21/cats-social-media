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
	cat.DELETE("/:id", handler.DeleteCat)

	cat.POST("/match", handler.MatchCat)
	cat.DELETE("/match/:id", handler.DeleteMatchCat)
	cat.POST("/match/reject", handler.RejectMatchCat)
	cat.GET("/match", handler.GetListMatchCat)
	cat.POST("/match/approve", handler.MatchApprove)

}
