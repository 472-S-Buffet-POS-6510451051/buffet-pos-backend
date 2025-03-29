package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID 			uuid.UUID 	`json:"id" gorm:"primaryKey;type:uuid"` 
	Phone 		string 		`json:"phone" gorm:"type:varchar(10)"`
	PIN 		string 		`json:"pin" gorm:"varchar(255)"`
	Point		int 		`json:"point" gorm:"integer"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	DeletedAt 	time.Time	`json:"deletedAt"`
}