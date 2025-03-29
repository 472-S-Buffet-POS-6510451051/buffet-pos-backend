package requests

type UpdateInvoiceStatusRequest struct {
	InvoiceID string `json:"invoice_id" validate:"required"`
}

type ChargeFeeFoodOverWeightRequest struct {
	InvoiceID 		string 	`json:"invoice_id" validate:"required"`
	TotalFoodWeight float64 `json:"total_food_weight" validate:"required"`
}
