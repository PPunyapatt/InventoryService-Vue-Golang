package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/pkg/errors"
)

func (i *inventoryRepository) ListInventories(ctx context.Context, pagination *helper.Pagination) ([]*constant.Inventories, error) {
	query := `
		SELECT
			inventory_code, created_at
		FROM inventories
		ORDER BY created_at
	`

	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM (%s)`, query)
	err := i.sqlx.QueryRowxContext(
		ctx,
		countQuery,
	).Scan(&pagination.TotalCount)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		return nil, err
	}
	args := []interface{}{pagination.Offset, pagination.Limit}
	query += fmt.Sprintf(` OFFSET $1 LIMIT $2`)

	inventoryList := []*constant.Inventories{}
	err = i.sqlx.SelectContext(ctx, &inventoryList, query, args...)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		return nil, err
	}

	return inventoryList, nil
}

func (i *inventoryRepository) AddInventory(ctx context.Context, inventoryCode string) error {
	query := `
		INSERT INTO inventories (
			inventory_code
		)
		VALUES ($1)
	`
	args := []interface{}{inventoryCode}
	_, err := i.sqlx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		return err
	}

	return nil
}

func (i *inventoryRepository) DeleteInventory(ctx context.Context, inventoryCode string) error {
	query := `
		DELETE 
		FROM inventories
		WHERE inventory_code = $1
	`

	agrs := []interface{}{inventoryCode}
	_, err := i.sqlx.ExecContext(ctx, query, agrs...)
	if err != nil {
		log.Printf("%+v", errors.WithStack(err))
		return err
	}
	return nil
}
