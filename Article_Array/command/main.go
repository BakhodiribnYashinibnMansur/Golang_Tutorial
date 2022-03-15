package main

import (
	"article/api"
	"article/api/handlers"
	"article/configs"
	"article/storage/postgres"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()
	pgStore := postgres.NewPostgres(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	))

	h := handlers.NewHandler(pgStore)

	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api.SetUpAPI(r, h, cfg)

	r.Run(":8080")
}
