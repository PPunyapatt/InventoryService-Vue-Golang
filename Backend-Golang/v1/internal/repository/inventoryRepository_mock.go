package repository

import (
	"context"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/stretchr/testify/mock"
)

type MockInventoryRepo struct {
	mock.Mock
}

func NewInventoryRepoMock() *MockInventoryRepo {
	return &MockInventoryRepo{}
}

func (m *MockInventoryRepo) ListInventories(ctx context.Context, p *helper.Pagination) ([]*constant.Inventories, error) {
	args := m.Called(ctx, p)
	return args.Get(0).([]*constant.Inventories), args.Error(1)
}

func (m *MockInventoryRepo) AddInventory(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}

func (m *MockInventoryRepo) DeleteInventory(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
