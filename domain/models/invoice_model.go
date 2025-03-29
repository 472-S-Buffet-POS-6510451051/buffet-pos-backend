package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID          	 uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid"`
	PeopleAmount	 int        `json:"peopleAmount" gorm:"type:integer"`
	TotalPrice   	 float64    `json:"totalPrice" gorm:"type:decimal"`
	IsPaid      	 bool       `json:"isPaid" gorm:"type:boolean"`
	CustomerPhone	 string		`json:"customer_phone" gorm:"type:varchar(100)"`
	PriceFeeFoodOverWeight float64 `json:"price_fee_food_overweight" gorm:"type:decimal"`
	TableID    	 	 *uuid.UUID `json:"tableId" gorm:"type:uuid;foreignKey:TableID"`
	CreatedAt    	 time.Time  `json:"createdAt"`
	UpdatedAt    	 time.Time  `json:"updatedAt"`
	DeletedAt    	 gorm.DeletedAt
}
