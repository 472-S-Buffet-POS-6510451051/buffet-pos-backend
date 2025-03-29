package requests

type OrderItemRequest struct {
	MenuID   string `json:"menu_id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}