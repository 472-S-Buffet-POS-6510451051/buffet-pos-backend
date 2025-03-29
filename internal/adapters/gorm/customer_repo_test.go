package gorm

import (
	"context"
	"testing"
	"time"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the Customer repository
type MockCustomerRepo struct {
	mock.Mock
}

func (m *MockCustomerRepo) Create(ctx context.Context, req *requests.CustomerRegisterRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockCustomerRepo) FindAll(ctx context.Context) ([]responses.BaseCustomer, error) {
	args := m.Called(ctx)
	return args.Get(0).([]responses.BaseCustomer), args.Error(1)
}

func (m *MockCustomerRepo) FindByID(ctx context.Context, customerID string) (*models.Customer, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepo) FindByPhone(ctx context.Context, phone string) (*models.Customer, error) {
	args := m.Called(ctx, phone)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepo) AddPoint(ctx context.Context, req *requests.CustomerAddPointRequest) (*responses.BaseCustomer, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*responses.BaseCustomer), args.Error(1)
}

func (m *MockCustomerRepo) RedeemPoint(ctx context.Context, req *requests.CustomerRedeemRequest, usePoint int) (*responses.BaseCustomer, error) {
	args := m.Called(ctx, req, usePoint)
	return args.Get(0).(*responses.BaseCustomer), args.Error(1)
}

func (m *MockCustomerRepo) Delete(ctx context.Context, customerID string) error {
	args := m.Called(ctx, customerID)
	return args.Error(0)
}

// Test for Create function
func TestCreateCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	req := &requests.CustomerRegisterRequest{
		Phone: "1234567890",
		PIN:   "123456",
	}

	mockRepo.On("Create", mock.Anything, req).Return(nil)

	err := mockRepo.Create(context.Background(), req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test for FindAll function
func TestFindAllCustomers(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	expectedCustomers := []responses.BaseCustomer{
		{ID: uuid.New(), Phone: "1234567890", Point: 10},
		{ID: uuid.New(), Phone: "0987654321", Point: 10},
	}

	mockRepo.On("FindAll", mock.Anything).Return(expectedCustomers, nil)

	customers, err := mockRepo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)
	mockRepo.AssertExpectations(t)
}

// Test for FindByID function
func TestFindByID(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	customerID := uuid.New()
	expectedCustomer := &models.Customer{
		ID:        customerID,
		Phone:     "1234567890",
		PIN:       "123456",
		Point:     10,
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, customerID.String()).Return(expectedCustomer, nil)

	customer, err := mockRepo.FindByID(context.Background(), customerID.String())
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)
	mockRepo.AssertExpectations(t)
}

// Test for FindByPhone function
func TestFindByPhone(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	expectedCustomer := &models.Customer{
		ID:    uuid.New(),
		Phone: "1234567890",
		PIN:   "123456",
		Point: 10,
	}

	mockRepo.On("FindByPhone", mock.Anything, expectedCustomer.Phone).Return(expectedCustomer, nil)

	customer, err := mockRepo.FindByPhone(context.Background(), expectedCustomer.Phone)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)
	mockRepo.AssertExpectations(t)
}

// Test for AddPoint function
func TestAddPoint(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	req := &requests.CustomerAddPointRequest{
		Phone: "1234567890",
		PIN:   "123456",
		Point: 10,
	}
	expectedCustomer := &responses.BaseCustomer{
		ID:    uuid.New(),
		Phone: req.Phone,
		Point: 10,
	}

	mockRepo.On("AddPoint", mock.Anything, req).Return(expectedCustomer, nil)

	customer, err := mockRepo.AddPoint(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)
	mockRepo.AssertExpectations(t)
}

// Test for RedeemPoint function
func TestRedeemPoint(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	req := &requests.CustomerRedeemRequest{
		Phone:     "1234567890",
		PIN:       "123456",
		InvoiceID: "INV12345",
	}
	expectedCustomer := &responses.BaseCustomer{
		ID:    uuid.New(),
		Phone: req.Phone,
		Point: 10,
	}

	mockRepo.On("RedeemPoint", mock.Anything, req, 50).Return(expectedCustomer, nil)

	customer, err := mockRepo.RedeemPoint(context.Background(), req, 50)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)
	mockRepo.AssertExpectations(t)
}

// Test for Delete function
func TestDeleteCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	customerID := uuid.New().String()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := mockRepo.Delete(context.Background(), customerID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
