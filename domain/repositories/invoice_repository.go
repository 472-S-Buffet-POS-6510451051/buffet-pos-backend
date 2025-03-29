package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
)

type InvoiceRepository interface {
	FindByID(ctx context.Context, invoiceID string) (*models.Invoice, error)
	Create(ctx context.Context, tableID string, totalPrice float64, peopleAmount int) error
	SetPaid(ctx context.Context, invoiceID string) error
	Cancel(ctx context.Context, invoiceID string) error
	GetAllUnpaid(ctx context.Context) ([]models.Invoice, error)
	GetAllPaid(ctx context.Context) ([]models.Invoice, error)
	GetByTableID(ctx context.Context, tableID string) (*models.Invoice, error)
	AddTotalPriceByID(ctx context.Context, invoiceID string, price float64) error
	SetPriceFeeFoodOverWeightByID(ctx context.Context, invoiceID string, price float64) error
	SetCustomerPhoneByID(ctx context.Context, invoiceID string, customerPhone string) error
}
