package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type OrderItemRepository interface {
	GetOrderItemsByOrderID(ctx context.Context, orderID string) ([]models.OrderItem, error)
	CreateOrderItem(ctx context.Context, orderItem *requests.OrderItemRequest, orderID string) error
	GetAmountBestSellingMenu(ctx context.Context, amount int) ([]responses.NumberMenu, error)
}
