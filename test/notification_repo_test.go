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

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) Create(ctx context.Context, req *requests.StaffNotificationRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockNotificationRepository) FindAll(ctx context.Context) ([]*models.StaffNotification, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.StaffNotification), args.Error(1)
}

func (m *MockNotificationRepository) FindByID(ctx context.Context) (*models.StaffNotification, error) {
	args := m.Called(ctx)
	return args.Get(0).(*models.StaffNotification), args.Error(1)
}

func (m *MockNotificationRepository) Update(ctx context.Context, req *requests.StaffNotificationRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockNotificationRepository) Delete(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestCreateNotification(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	req := &requests.StaffNotificationRequest{
		TableID: uuid.New(),
	}

	mockRepo.On("Create", context.Background(), req).Return(nil)

	err := mockRepo.Create(context.Background(), req)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestShowListOfNotification(t *testing.T) {
	mockRepo := new(MockNotificationRepository)

	mockNotification := &models.StaffNotification{
		ID:      uuid.New(),
		TableID: uuid.New(),
	}

	mockRepo.On("FindAll", context.Background()).Return([]*models.StaffNotification{mockNotification}, nil)

	result, err := mockRepo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, mockNotification, result[0])

	mockRepo.AssertExpectations(t)
}

func TestUpdateNotification(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	req := &requests.StaffNotificationRequest{
		TableID: uuid.New(),
	}

	mockRepo.On("Update", context.Background(), req).Return(nil)

	err := mockRepo.Update(context.Background(), req)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteNotification(t *testing.T) {
	mockRepo := new(MockNotificationRepository)

	mockRepo.On("Delete", context.Background()).Return(nil)

	err := mockRepo.Delete(context.Background())
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestShowNotificationByID(t *testing.T) {
	mockRepo := new(MockNotificationRepository)

	mockNotification := &models.StaffNotification{
		ID:      uuid.New(),
		TableID: uuid.New(),
	}

	mockRepo.On("FindByID", context.Background()).Return(mockNotification, nil)

	result, err := mockRepo.FindByID(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, mockNotification, result)

	mockRepo.AssertExpectations(t)
}
