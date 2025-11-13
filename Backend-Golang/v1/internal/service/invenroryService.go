package service

import (
	"context"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
)

func (i *inventoryService) ListInventories(ctx context.Context, pagination *helper.Pagination) ([]*constant.Inventories, error) {
	// time.Sleep(time.Second * 15)
	data, err := i.inventoryRepo.ListInventories(ctx, pagination)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i *inventoryService) AddInventory(ctx context.Context, inventoryCode string) error {
	if err := i.inventoryRepo.AddInventory(ctx, inventoryCode); err != nil {
		return err
	}
	return nil
}

func (i *inventoryService) DeleteInventory(ctx context.Context, inventoryCode string) error {
	if err := i.inventoryRepo.DeleteInventory(ctx, inventoryCode); err != nil {
		return err
	}
	return nil
}
