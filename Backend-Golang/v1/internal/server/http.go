package server

import (
	"context"
	"log"
	"log/slog"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/api"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MapRoutes(ctx context.Context, handler *handler.ApiHandler) {
	app := fiber.New()
	errCh := make(chan error, 1)
	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})

	// Add logger middleware
	app.Use(c)
	app.Use(logger.New())

	// routes
	api.Route(app, handler)

	go func() {
		err := app.Listen(":1234")
		if err != nil {
			slog.Error("❌ Fiber error", "error", err.Error())
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("🛑 context done, shutting down fiber")
		if err := app.Shutdown(); err != nil {
			slog.Error("❌ Fiber shutdown error", "error", err)
		}
	case err := <-errCh:
		slog.Error("❌ Fiber error", "error", err.Error())
	}
}
