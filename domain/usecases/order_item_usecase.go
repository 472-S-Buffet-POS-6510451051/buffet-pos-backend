package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type OrderItemUseCase interface {
	GetBestSellingMenu(ctx context.Context) ([]responses.NumberMenu, error)
}

type orderItemService struct {
	orderItemRepo repositories.OrderItemRepository
	config        *configs.Config
}

func NewOrderItemService(orderItemRepo repositories.OrderItemRepository, config *configs.Config) OrderItemUseCase {
	return &orderItemService{
		orderItemRepo: orderItemRepo,
		config:        config,
	}
}

func (o *orderItemService) GetBestSellingMenu(ctx context.Context) ([]responses.NumberMenu, error) {
	menus, err := o.orderItemRepo.GetAmountBestSellingMenu(ctx, 6)
	if err != nil {
		return nil, err
	}

	if menus == nil {
		return []responses.NumberMenu{}, nil
	}

	return menus, nil
}
