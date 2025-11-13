package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/constant"
	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (c *ApiHandler) ListInventories(ctx *fiber.Ctx) error {
	context_, cancel := context.WithTimeout(ctx.Context(), time.Second*10)
	defer cancel()
	pagination := helper.PaginationNew(ctx)

	data, err := c.InventoryService.ListInventories(context_, pagination)
	if err != nil {
		return helper.RespondHttpError(ctx, helper.NewHttpError(http.StatusBadRequest, err))
	}

	return ctx.Status(http.StatusOK).JSON(constant.ListResponse{
		StatusCode: http.StatusOK,
		Data:       data,
		Message:    "Data retrieved",
		Pagination: pagination,
	})
}

func (c *ApiHandler) AddInventory(ctx *fiber.Ctx) error {
	context_, cancel := context.WithTimeout(ctx.Context(), time.Second*10)
	defer cancel()

	request, err := helper.ParseAndValidateRequest(ctx, &constant.Inventories{})
	if err != nil {
		return helper.RespondHttpError(ctx, helper.NewHttpError(http.StatusBadRequest, err))
	}
	fmt.Println("request: ", request)

	err = c.InventoryService.AddInventory(context_, request.InventoryCode)
	if err != nil {
		return helper.RespondHttpError(ctx, helper.NewHttpError(http.StatusInternalServerError, err))
	}

	return ctx.Status(http.StatusOK).JSON(constant.StatusResponse{
		StatusCode: http.StatusOK,
		Message:    "Add inventory success",
	})
}

func (c *ApiHandler) DeleteInventory(ctx *fiber.Ctx) error {
	context_, cancel := context.WithTimeout(ctx.Context(), time.Second*10)
	defer cancel()

	request, err := helper.ParseAndValidateRequest(ctx, &constant.Inventories{})
	if err != nil {
		return helper.RespondHttpError(ctx, helper.NewHttpError(http.StatusBadRequest, err))
	}

	err = c.InventoryService.DeleteInventory(context_, request.InventoryCode)
	if err != nil {
		return helper.RespondHttpError(ctx, helper.NewHttpError(http.StatusInternalServerError, err))
	}

	return ctx.Status(http.StatusOK).JSON(constant.StatusResponse{
		StatusCode: http.StatusOK,
		Message:    "Delete inventory success",
	})
}
