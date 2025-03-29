package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StaffNotificationStatus string

const (
	PENDING		StaffNotificationStatus = "pending"
	ACCEPTED  	StaffNotificationStatus = "accepted"
	REJECTED 	StaffNotificationStatus = "rejected"
)

type StaffNotification struct {
	ID			uuid.UUID					`json:"id" gorm:"primaryKey;type:uuid"`
	TableID 	uuid.UUID					`json:"table_id" gorm:"foreignKey:TableID;constraint:OnDelete:CASCADE;type:uuid"`
	Status  	StaffNotificationStatus 	`json:"status" gorm:"type:varchar(50);default:'pending'"`
	CreatedAt 	time.Time 					`json:"createdAt"`
	UpdatedAt  	time.Time   				`json:"updatedAt"`
	DeletedAt   gorm.DeletedAt
}