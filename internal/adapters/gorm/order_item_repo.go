package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItemGormRepository struct {
	DB *gorm.DB
}

func NewOrderItemGormRepository(db *gorm.DB) *OrderItemGormRepository {
	return &OrderItemGormRepository{
		DB: db,
	}
}

func (o *OrderItemGormRepository) GetOrderItemsByOrderID(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	result := o.DB.Where("order_id = ?", orderID).Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}

func (o *OrderItemGormRepository) CreateOrderItem(ctx context.Context, orderItem *requests.OrderItemRequest, orderID string) error {
	newOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return err
	}

	menuID, err := uuid.Parse(orderItem.MenuID)
	if err != nil {
		return err
	}

	id := uuid.New()

	newOrderItem := &models.OrderItem{
		ID:       id,
		OrderID:  newOrderID,
		MenuID:   menuID,
		Quantity: orderItem.Quantity,
	}
	result := o.DB.Create(newOrderItem)
	return result.Error
}

func (o *OrderItemGormRepository) GetAmountBestSellingMenu(ctx context.Context, amount int) ([]responses.NumberMenu, error) {
	var numberMenus []responses.NumberMenu

	err := o.DB.Table("order_items").
		Select("menus.id, menus.name, menus.description, menus.category_id, menus.image_url, menus.is_available, SUM(order_items.quantity) as total_quantity").
		Joins("JOIN menus ON menus.id = order_items.menu_id").
		Group("menus.id, menus.name, menus.description, menus.category_id, menus.image_url, menus.is_available").
		Order("total_quantity DESC").
		Limit(amount).
		Scan(&numberMenus).Error

	if err != nil {
		return nil, err
	}

	// เพิ่มหมายเลขอันดับกำกับให้กับแต่ละเมนู
	for i := range numberMenus {
		numberMenus[i].Number = i + 1
	}

	return numberMenus, nil
}