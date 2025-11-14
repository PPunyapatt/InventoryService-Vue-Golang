package handler_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/api/handler"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddInventoryHandler(t *testing.T) {
	t.Run("Add inventory handler success", func(t *testing.T) {
		app := fiber.New()
		mockSvc := service.NewInventoryServiceMock()
		h := handler.ServiceNew(mockSvc)

		app.Post("/inventory", h.AddInventory)

		body := `{"inventory_code": "TEST123"}`

		mockSvc.On("AddInventory", mock.Anything, "TEST123").
			Return(nil)

		req := httptest.NewRequest("POST", "/inventory", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Add inventory handler failed", func(t *testing.T) {
		app := fiber.New()
		mockSvc := service.NewInventoryServiceMock()
		h := handler.ServiceNew(mockSvc)

		app.Post("/inventory", h.AddInventory)

		body := `{"inventory_code": "TEST123"}`

		mockSvc.On("AddInventory", mock.Anything, "TEST123").
			Return(&pgconn.PgError{
				Code:    "23505", // duplicate key
				Message: "duplicate key value violates unique constraint",
			})

		req := httptest.NewRequest("POST", "/inventory", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 409, resp.StatusCode)
	})

}

func TestListInventories(t *testing.T) {
	t.Run("List inventories handler", func(t *testing.T) {
		app := fiber.New()
		mockSvc := service.NewInventoryServiceMock()
		h := handler.ServiceNew(mockSvc)

		app.Get("/inventories", h.ListInventories)

		p := &helper.Pagination{
			Page:  1,
			Limit: 10,
		}
		mockResponse := []*constant.Inventories{
			{InventoryCode: "A"},
			{InventoryCode: "B"},
		}

		mockSvc.On("ListInventories", mock.Anything, p).
			Return(mockResponse, nil)

		req := httptest.NewRequest("GET", "/inventories", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

func TestDeleteInventory(t *testing.T) {
	t.Run("List inventories handler success", func(t *testing.T) {
		app := fiber.New()
		mockSvc := service.NewInventoryServiceMock()
		h := handler.ServiceNew(mockSvc)

		app.Delete("/inventory", h.ListInventories)

		body := `{"inventory_code": "TEST123"}`

		mockSvc.On("DeleteInventory", mock.Anything, "TEST123").
			Return(nil)

		req := httptest.NewRequest("DELETE", "/inventory", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
