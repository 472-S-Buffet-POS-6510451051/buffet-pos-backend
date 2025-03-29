package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuGormRepository struct {
	DB *gorm.DB
}

func NewMenuGormRepository(db *gorm.DB) *MenuGormRepository {
	return &MenuGormRepository{
		DB: db,
	}
}

func (m *MenuGormRepository) Create(ctx context.Context, req *requests.AddMenuRequest, imageURL string) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return err
	}

	menu := &models.Menu{
		ID:          id,
		Name:        req.Name,
		Description: &req.Description,
		ImageURL:    &imageURL,
		CategoryID:  &categoryID,
	}

	result := m.DB.Create(menu)

	return result.Error
}

func (m *MenuGormRepository) FindAll(ctx context.Context) ([]models.Menu, error) {
	var menus []models.Menu
	result := m.DB.Find(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}

func (m *MenuGormRepository) FindByID(ctx context.Context, menuID string) (*models.Menu, error) {
	var menu models.Menu
	result := m.DB.Where("id = ?", menuID).First(&menu)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil
}

func (m *MenuGormRepository) Edit(ctx context.Context, req *requests.EditMenuRequest, imageURL string) error {
	var menu models.Menu
	if err := m.DB.Where("id = ?", req.ID).First(&menu).Error; err != nil {
		return err
	}

	categoryID, _ := uuid.Parse(req.CategoryID)
	if categoryID == uuid.Nil && menu.CategoryID != nil {
		categoryID = *menu.CategoryID
	}

	if req.Name != "" {
		menu.Name = req.Name
	}
	if req.Description != "" {
		menu.Description = &req.Description
	}
	if imageURL != "" {
		menu.ImageURL = &imageURL
	}
	if req.CategoryID != "" {
		menu.CategoryID = &categoryID
	}
	menu.IsAvailable = req.IsAvailable

	result := m.DB.Save(&menu)
	return result.Error
}

func (m *MenuGormRepository) Delete(ctx context.Context, menuID string) error {
	result := m.DB.Where("id = ?", menuID).Delete(&models.Menu{})
	return result.Error
}

func (m *MenuGormRepository) FindByName(ctx context.Context, name string) (*models.Menu, error) {	
	var menu models.Menu
	result := m.DB.Where("name = ?", name).First(&menu)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil
}
