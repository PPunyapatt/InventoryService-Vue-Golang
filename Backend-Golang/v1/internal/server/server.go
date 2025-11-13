package server

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/api/handler"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/pkg/config"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/pkg/database"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/repository"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/service"
	"github.com/joho/godotenv"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, skipping...")
		return err
	}

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

	repository := repository.NewOrderRepository(db.Sqlx)

	service := service.NewInventoryService(repository)

	handler := handler.ServiceNew(service)

	MapRoutes(ctx, handler)

	return nil
}
