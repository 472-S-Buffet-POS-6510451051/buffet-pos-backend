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

type MockMenuRepository struct {
	mock.Mock
}

func (m *MockMenuRepository) Create(ctx context.Context, req *requests.AddMenuRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockMenuRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Menu, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Menu), args.Error(1)
}

func (m *MockMenuRepository) FindAll(ctx context.Context) ([]*models.Menu, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Menu), args.Error(1)
}

func (m *MockMenuRepository) Update(ctx context.Context, req *requests.AddMenuRequest, id uuid.UUID) error {
	args := m.Called(ctx, req, id)
	return args.Error(0)
}

func (m *MockMenuRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMenuRepository) GetBestSellerMenu(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateMenu(t *testing.T) {
	mockRepo := new(MockMenuRepository)
	req := &requests.AddMenuRequest{
		Name:        "test",
		Description: "test",
		CategoryID:  "2",
	}

	mockRepo.On("Create", context.Background(), req).Return(nil)

	err := mockRepo.Create(context.Background(), req)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetBestSellerMenu(t *testing.T) {
	mockRepo := new(MockMenuRepository)

	mockRepo.On("GetBestSellerMenu", context.Background(), uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")).Return(nil)

	err := mockRepo.GetBestSellerMenu(context.Background(), uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockMenuRepository)

	mockRepo.On("Delete", context.Background(), uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")).Return(nil)

	err := mockRepo.Delete(context.Background(), uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockMenuRepository)
	req := &requests.AddMenuRequest{
		Name:        "test",
		Description: "test",
		CategoryID:  "2",
	}

	mockRepo.On("Update", context.Background(), req, uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")).Return(nil)

	err := mockRepo.Update(context.Background(), req, uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestFindAllMenu(t *testing.T) {
	mockRepo := new(MockMenuRepository)

	mockMenu := &models.Menu{
		ID:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
		Name: "test",
	}

	mockRepo.On("FindAll", context.Background()).Return([]*models.Menu{mockMenu}, nil)

	menu, err := mockRepo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, []*models.Menu{mockMenu}, menu)

	mockRepo.AssertExpectations(t)
}

func TestFindByIDMenu(t *testing.T) {
	mockRepo := new(MockMenuRepository)

	mockMenu := &models.Menu{
		ID:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
		Name: "test",
	}

	mockRepo.On("FindByID", context.Background(), mockMenu.ID).Return(mockMenu, nil)

	menu, err := mockRepo.FindByID(context.Background(), mockMenu.ID)
	assert.NoError(t, err)
	assert.Equal(t, mockMenu, menu)

	mockRepo.AssertExpectations(t)
}