package requests

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/google/uuid"
)

type StaffNotificationRequest struct {
	TableID 	uuid.UUID	`json:"table_id" validate:"required"`
}

type UpdateStaffNotificationRequest struct {
	StaffNotificationID string							`json:"staff_notification_id"`
	Status 				models.StaffNotificationStatus	`json:"status"`
}