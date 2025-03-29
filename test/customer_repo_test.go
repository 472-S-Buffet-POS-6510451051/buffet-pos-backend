package test

import (
	"context"
	"testing"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) Create(ctx context.Context, req *requests.CustomerRegisterRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockCustomerRepository) FindByPhone(ctx context.Context, phone string) (*models.Customer, error) {
	args := m.Called(ctx, phone)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) AddPoint(ctx context.Context, phone string, point int) error {
	args := m.Called(ctx, phone, point)
	return args.Error(0)
}

func (m *MockCustomerRepository) RedeemPoint(ctx context.Context, phone string, point int) error {
	args := m.Called(ctx, phone, point)
	return args.Error(0)
}

func TestCreateCustomerWithPhoneAndPin(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	req := &requests.CustomerRegisterRequest{
		Phone: "0123456789",
		PIN:   "1234",
	}

	mockRepo.On("Create", context.Background(), req).Return(nil)

	err := mockRepo.Create(context.Background(), req)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestShowListOfCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepository)

	mockCustomer := &models.Customer{
		ID:    uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
		Phone: "0123456789",
		Point: 0,
	}

	mockRepo.On("FindByPhone", context.Background(), mockCustomer.Phone).Return(mockCustomer, nil)

	customer, err := mockRepo.FindByPhone(context.Background(), mockCustomer.Phone)
	assert.NoError(t, err)
	assert.Equal(t, mockCustomer, customer)

	mockRepo.AssertExpectations(t)
}

func TestAddpointToCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepository)

	mockCustomer := &models.Customer{
		Phone: "0123456789",
	}

	mockRepo.On("AddPoint", context.Background(), mockCustomer.Phone, 10).Return(nil)

	err := mockRepo.AddPoint(context.Background(), mockCustomer.Phone, 10)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestRedeemPointFromCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepository)

	mockCustomer := &models.Customer{
		Phone: "0123456789",
	}

	mockRepo.On("RedeemPoint", context.Background(), mockCustomer.Phone, 10).Return(nil)

	err := mockRepo.RedeemPoint(context.Background(), mockCustomer.Phone, 10)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
