package rest

import (
    "github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
    "github.com/gofiber/fiber/v2"
)

type OrderItemHandler interface {
    GetBestSellingMenu(c *fiber.Ctx) error
}

type orderItemHandler struct {
    service usecases.OrderItemUseCase
}

func NewOrderItemHandler(service usecases.OrderItemUseCase) OrderItemHandler {
    return &orderItemHandler{
        service: service,
    }
}

// Get Best Selling Menu
// @Summary Get Best Selling Menu
// @Description Get Best Selling Menu
// @Tags General
// @Accept json
// @Produce json
// @Success 200 {array} responses.NumberMenu
// @Router /general/menus/best-selling [get]
// @Security ApiKeyAuth
func (o *orderItemHandler) GetBestSellingMenu(c *fiber.Ctx) error {
    menus, err := o.service.GetBestSellingMenu(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(menus)
}