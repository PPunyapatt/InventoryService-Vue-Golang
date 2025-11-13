package repository

import (
	"context"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/jmoiron/sqlx"
)

type inventoryRepository struct {
	sqlx *sqlx.DB
}

type InventoryRepository interface {
	ListInventories(ctx context.Context, pagination *helper.Pagination) ([]*constant.Inventories, error)
	AddInventory(ctx context.Context, inventory_code string) error
	DeleteInventory(ctx context.Context, inventory_code string) error
}

func NewOrderRepository(sqlx *sqlx.DB) InventoryRepository {
	return &inventoryRepository{
		sqlx: sqlx,
	}
}
