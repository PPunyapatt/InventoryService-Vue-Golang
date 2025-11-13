package api

import (
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(
	app *fiber.App,
	api *handler.ApiHandler,
) {
	version := "/api/v1"

	inventory := app.Group(version + "/")
	inventory.Get("/inventories", api.ListInventories)
	inventory.Delete("/inventory", api.DeleteInventory)
	inventory.Post("/inventory", api.AddInventory)
}
