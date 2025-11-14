package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/repository"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/service"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListInventories(t *testing.T) {
	t.Run("List inventories service success", func(t *testing.T) {
		mockRepo := repository.NewInventoryRepoMock()
		svc := service.NewInventoryService(mockRepo)

		p := &helper.Pagination{
			Page:  1,
			Limit: 10,
		}

		mockData := []*constant.Inventories{
			{InventoryCode: "A"},
			{InventoryCode: "B"},
		}

		mockRepo.On("ListInventories", mock.Anything, p).Return(mockData, nil)

		data, err := svc.ListInventories(context.Background(), p)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(data))
	})

}

func TestDeleteInventory(t *testing.T) {
	t.Run("Delete inventories service failed", func(t *testing.T) {
		mockRepo := repository.NewInventoryRepoMock()
		svc := service.NewInventoryService(mockRepo)

		mockRepo.On("DeleteInventory", mock.Anything, "ABC").
			Return(errors.New("delete failed"))

		err := svc.DeleteInventory(context.Background(), "ABC")
		assert.Error(t, err)
	})

	t.Run("Delete inventories service success", func(t *testing.T) {
		mockRepo := repository.NewInventoryRepoMock()
		svc := service.NewInventoryService(mockRepo)

		mockRepo.On("DeleteInventory", mock.Anything, "ABB").
			Return(nil)
		err := svc.DeleteInventory(context.Background(), "ABB")
		assert.NoError(t, err)
	})

}

func TestAddInventory(t *testing.T) {
	t.Run("Add inventory service failed", func(t *testing.T) {
		var pgError *pgconn.PgError
		mockRepo := repository.NewInventoryRepoMock()
		svc := service.NewInventoryService(mockRepo)

		mockRepo.On("AddInventory", mock.Anything, "ABB").
			Return(&pgconn.PgError{
				Code:    "23505", // duplicate key
				Message: "duplicate key value violates unique constraint",
			})

		err := svc.AddInventory(context.Background(), "ABB")
		assert.Error(t, err)
		assert.ErrorAs(t, err, &pgError)
		assert.Equal(t, "23505", pgError.Code)
	})

	t.Run("Add inventory service success", func(t *testing.T) {
		mockRepo := repository.NewInventoryRepoMock()
		svc := service.NewInventoryService(mockRepo)

		mockRepo.On("AddInventory", mock.Anything, "ABB").
			Return(nil)

		err := svc.AddInventory(context.Background(), "ABB")
		assert.NoError(t, err)
	})

}
