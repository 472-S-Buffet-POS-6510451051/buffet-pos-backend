package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StaffNotificationGormRepository struct {
	DB *gorm.DB
}

func NewStaffNotificationGormRepository(db *gorm.DB) repositories.StaffNotificationRepository {
	return &StaffNotificationGormRepository{
		DB: db,
	}
}

func (s *StaffNotificationGormRepository) Create(ctx context.Context, req *requests.StaffNotificationRequest) error {
	// Generate UUID
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	staffNotification := &models.StaffNotification{
		ID:      id,
		TableID: req.TableID,
		Status:  models.PENDING,
	}

	return s.DB.Create(staffNotification).Error
}

func (s *StaffNotificationGormRepository) FindByID(ctx context.Context, staffNotificationID string) (*responses.StaffNotificationBase, error) {
	var notification models.StaffNotification
	if err := s.DB.First(&notification, "id = ?", staffNotificationID).Error; err != nil {
		return nil, err
	}

	return &responses.StaffNotificationBase{
		ID:        notification.ID.String(),
		TableID:   notification.TableID.String(),
		Status:    notification.Status,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}, nil
}

func (s *StaffNotificationGormRepository) FindByTableID(ctx context.Context, tableID string) (*responses.StaffNotificationBase, error) {
	var notification models.StaffNotification
	if err := s.DB.Last(&notification, "table_id = ?", tableID).Error; err != nil {
		return nil, err
	}

	return &responses.StaffNotificationBase{
		ID:        notification.ID.String(),
		TableID:   notification.TableID.String(),
		Status:    notification.Status,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}, nil
}

func (s *StaffNotificationGormRepository) GetByTableID(ctx context.Context, tableID string) (*responses.StaffNotificationBase, error) {
	var notification models.StaffNotification
	if err := s.DB.Order("createdAt desc").First(&notification, "table_id = ?", tableID).Error; err != nil {
		return nil, err
	}

	return &responses.StaffNotificationBase{
		ID:        notification.ID.String(),
		TableID:   notification.TableID.String(),
		Status:    notification.Status,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}, nil
}

func (s *StaffNotificationGormRepository) GetAll(ctx context.Context) ([]responses.StaffNotificationBase, error) {
	var notifications []models.StaffNotification
	if err := s.DB.Find(&notifications).Error; err != nil {
		return nil, err
	}

	var response []responses.StaffNotificationBase
	for _, notification := range notifications {
		response = append(response, responses.StaffNotificationBase{
			ID:        notification.ID.String(),
			TableID:   notification.TableID.String(),
			Status:    notification.Status,
			CreatedAt: notification.CreatedAt,
			UpdatedAt: notification.UpdatedAt,
		})
	}

	return response, nil
}

func (s *StaffNotificationGormRepository) GetAllByStatus(ctx context.Context, status string) ([]responses.StaffNotificationBase, error) {
	var notifications []models.StaffNotification
	if err := s.DB.WithContext(ctx).Where("status = ?", status).Find(&notifications).Error; err != nil {
		return nil, err
	}

	var response []responses.StaffNotificationBase
	for _, notification := range notifications {
		response = append(response, responses.StaffNotificationBase{
			ID:        notification.ID.String(),
			TableID:   notification.TableID.String(),
			Status:    notification.Status,
			CreatedAt: notification.CreatedAt,
			UpdatedAt: notification.UpdatedAt,
		})
	}

	return response, nil
}

func (s *StaffNotificationGormRepository) UpdateStatus(ctx context.Context, req *requests.UpdateStaffNotificationRequest) error {
	return s.DB.Model(&models.StaffNotification{}).
		Where("id = ?", req.StaffNotificationID).
		Update("status", req.Status).Error
}

func (s *StaffNotificationGormRepository) Delete(ctx context.Context, staffNotificationID string) error {
	return s.DB.Delete(&models.StaffNotification{}, "id = ?", staffNotificationID).Error
}
