package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type StaffNotificationRepository interface {
	Create(ctx context.Context, req *requests.StaffNotificationRequest) error
	GetAll(ctx context.Context) ([]responses.StaffNotificationBase, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStaffNotificationRequest) error
	GetAllByStatus(ctx context.Context, status string) ([]responses.StaffNotificationBase, error)

	FindByID(ctx context.Context, staffNotificationID string) (*responses.StaffNotificationBase, error)
	FindByTableID(ctx context.Context, tableID string) (*responses.StaffNotificationBase, error)
	Delete(ctx context.Context, staffNotificationID string) error
}