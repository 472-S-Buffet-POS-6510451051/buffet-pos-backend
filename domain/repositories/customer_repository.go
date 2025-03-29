package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type CustomerRepository interface {
	Create(ctx context.Context, req *requests.CustomerRegisterRequest) error
	FindAll(ctx context.Context) ([]responses.BaseCustomer, error)
	FindByID(ctx context.Context, customerID string) (*models.Customer, error)
	FindByPhone(ctx context.Context, phone string) (*models.Customer, error)
	AddPoint(ctx context.Context, req *requests.CustomerAddPointRequest) (*responses.BaseCustomer, error)
	RedeemPoint(ctx context.Context, req *requests.CustomerRedeemRequest, usePoint int, priceDiscount float64) (*responses.BaseCustomer, error)
	Delete(ctx context.Context, customerID string) error
}
