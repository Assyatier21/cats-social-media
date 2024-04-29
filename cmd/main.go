package main

import (
	"fmt"

	"github.com/backend-magang/cats-social-media/config"
	"github.com/backend-magang/cats-social-media/driver"
	"github.com/backend-magang/cats-social-media/internal/handler/api"
	"github.com/backend-magang/cats-social-media/internal/repository/postgres"
	"github.com/backend-magang/cats-social-media/internal/usecase"
	"github.com/backend-magang/cats-social-media/middleware"
	"github.com/backend-magang/cats-social-media/router"
	"github.com/labstack/echo/v4"
)

// @title           Swagger Backend Magang - Project 1
// @version         1.0
// @description     This is a documentation of Backend Magang - Project 1
func main() {
	server := echo.New()

	// Load Config
	cfg := config.Load()

	dbClient := driver.InitPostgres(cfg)

	postgresRepository := postgres.NewRepository(dbClient)
	usecase := usecase.NewUsecase(cfg, postgresRepository)
	handler := api.NewHandler(usecase)

	router.InitRouter(server, handler)
	middleware.InitMiddlewares(server)

	host := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	server.Start(host)
}
