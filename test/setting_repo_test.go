package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSettingRepository struct {
	mock.Mock
}

func (m *MockSettingRepository) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockSettingRepository) Set(ctx context.Context, key string, value string) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockSettingRepository) GetTransactionWithFeeFood(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func TestGetFeeFoodRate(t *testing.T) {
	mockRepo := new(MockSettingRepository)

	mockRepo.On("Get", context.Background(), "fee_food_rate").Return("10", nil)

	result, err := mockRepo.Get(context.Background(), "fee_food_rate")
	assert.NoError(t, err)
	assert.Equal(t, "10", result)

	mockRepo.AssertExpectations(t)
}

func TestSetFeeFoodRate(t *testing.T) {
	mockRepo := new(MockSettingRepository)

	mockRepo.On("Set", context.Background(), "fee_food_rate", "10").Return(nil)

	err := mockRepo.Set(context.Background(), "fee_food_rate", "10")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDisplayCalculateFeeFoodInTransaction(t *testing.T) {
	mockRepo := new(MockSettingRepository)

	mockRepo.On("GetTransactionWithFeeFood", context.Background()).Return("10", nil)

	result, err := mockRepo.GetTransactionWithFeeFood(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "10", result)

	mockRepo.AssertExpectations(t)
}

func TestSetCalculateFeeFoodInTransaction(t *testing.T) {
	mockRepo := new(MockSettingRepository)

	mockRepo.On("Set", context.Background(), "fee_food_rate", "10").Return(nil)

	err := mockRepo.Set(context.Background(), "fee_food_rate", "10")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}