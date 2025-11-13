package handler

import "github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/service"

type ApiHandler struct {
	InventoryService service.InventoryService
}

func ServiceNew(InventoryHandler service.InventoryService) *ApiHandler {
	return &ApiHandler{InventoryHandler}
}
