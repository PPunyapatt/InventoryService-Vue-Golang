package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/api/handler"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/pkg/config"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/pkg/database"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/pkg/redis"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/repository"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/service"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg, err := config.SetUpEnv()
	if err != nil {
		return err
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	db, err := database.InitDatabase(cfg)
	if err != nil {
		return err
	}
	defer db.Sqlx.Close()

	rdb, err := redis.InitRedis(cfg)
	if err != nil {
		return err
	}
	defer rdb.Close()

	repository := repository.NewInventoryRepository(db.Sqlx)

	service := service.NewInventoryService(repository)

	handler := handler.ServiceNew(service)

	MapRoutes(ctx, handler, rdb)

	return nil
}
