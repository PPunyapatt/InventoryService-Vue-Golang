package main

import (
	"log/slog"
	"os"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("server exited with error", "error", err)
		os.Exit(1)
	}
}
