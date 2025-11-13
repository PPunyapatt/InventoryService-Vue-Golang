package service

import (
	"context"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/repository"
)

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
}

type InventoryService interface {
	ListInventories(ctx context.Context, pagination *helper.Pagination) ([]*constant.Inventories, error)
	AddInventory(ctx context.Context, inventoryCode string) error
	DeleteInventory(ctx context.Context, inventoryCode string) error
}

func NewInventoryService(inventoryRepo repository.InventoryRepository) InventoryService {
	return &inventoryService{inventoryRepo}
}
