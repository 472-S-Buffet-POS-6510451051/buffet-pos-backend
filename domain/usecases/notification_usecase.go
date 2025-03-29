package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type StaffNotificationUseCase interface {
	NotifyStaff(ctx context.Context, req *requests.StaffNotificationRequest) error
	GetAllStaffNotification(ctx context.Context) ([]responses.StaffNotificationBase, error)
	GetAllStaffNotificationByStatus(ctx context.Context, status string) ([]responses.StaffNotificationBase, error)
	GetStaffNotificationByTableId(ctx context.Context, status string) (responses.StaffNotificationBase, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStaffNotificationRequest) error
}

type staffNotificationService struct {
	staffNotificationRepo repositories.StaffNotificationRepository
	config                *configs.Config
}

func NewStaffNotificationService(staffNotificationRepo repositories.StaffNotificationRepository, config *configs.Config) StaffNotificationUseCase {
	return &staffNotificationService{
		staffNotificationRepo: staffNotificationRepo,
		config:                config,
	}
}

func (s *staffNotificationService) GetAllStaffNotification(ctx context.Context) ([]responses.StaffNotificationBase, error) {
	notifications, err := s.staffNotificationRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}


func (s *staffNotificationService) GetAllStaffNotificationByStatus(ctx context.Context, status string) ([]responses.StaffNotificationBase, error) {

	validStatuses := map[models.StaffNotificationStatus]bool{
		models.PENDING:  true,
		models.ACCEPTED: true,
		models.REJECTED: true,
	}

	if !validStatuses[models.StaffNotificationStatus(status)] {
		return nil, exceptions.ErrorInvalidStaffNotificationStatus
	}
	
	return s.staffNotificationRepo.GetAllByStatus(ctx, status)
}

func (s *staffNotificationService) NotifyStaff(ctx context.Context, req *requests.StaffNotificationRequest) error {
	return s.staffNotificationRepo.Create(ctx, req)
}

func (s *staffNotificationService) UpdateStatus(ctx context.Context, req *requests.UpdateStaffNotificationRequest) error {
	validStatuses := map[models.StaffNotificationStatus]bool{
		models.PENDING:  true,
		models.ACCEPTED: true,
		models.REJECTED: true,
	}

	if !validStatuses[models.StaffNotificationStatus(req.Status)] {
		return exceptions.ErrorInvalidStaffNotificationStatus
	}

	if err := s.staffNotificationRepo.UpdateStatus(ctx, req); err != nil {
		return err
	}

	return nil
}

func (s* staffNotificationService) GetStaffNotificationByTableId(ctx context.Context, table_id string) (responses.StaffNotificationBase, error) {
	notification, err := s.staffNotificationRepo.FindByTableID(ctx, table_id)
	if err != nil {
		return responses.StaffNotificationBase{}, err
	}
	return *notification, nil
}