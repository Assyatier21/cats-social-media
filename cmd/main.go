package main

import (
	"fmt"

	"github.com/backendmagang/project-1/config"
	"github.com/labstack/echo/v4"
)

// @title           Swagger Backend Magang - Project 1
// @version         1.0
// @description     This is a documentation of Backend Magang - Project 1
func main() {
	server := echo.New()

	// Load Config
	cfg := config.Load()

	// dbClient := driver.InitPostgres(cfg.PostgresConfig)
	// esClient := driver.InitElasticClient(cfg.ElasticConfig)
	// redisClient := driver.InitRedisClient(cfg.RedisConfig)

	// postgresRepository := postgres.NewRepository(dbClient)
	// elasticRepository := elasticsearch.NewElasticRepository(esClient, cfg.ElasticConfig)

	// usecase := usecase.NewUsecase(postgresRepository, elasticRepository, redisClient)
	// handler := api.NewHandler(usecase)

	// router.InitRouter(server, handler)
	// middleware.InitMiddlewares(server)

	host := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	server.Start(host)
}
