package service

import (
	"context"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/stretchr/testify/mock"
)

type MockInventoryService struct {
	mock.Mock
}

func NewInventoryServiceMock() *MockInventoryService {
	return &MockInventoryService{}
}

func (m *MockInventoryService) ListInventories(ctx context.Context, p *helper.Pagination) ([]*constant.Inventories, error) {
	args := m.Called(ctx, p)
	return args.Get(0).([]*constant.Inventories), args.Error(1)
}

func (m *MockInventoryService) AddInventory(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}

func (m *MockInventoryService) DeleteInventory(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
