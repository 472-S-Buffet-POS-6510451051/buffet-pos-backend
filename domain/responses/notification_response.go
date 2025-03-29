package responses

import (
	"time"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
)

	type StaffNotificationBase struct {
		ID			string							`json:"id"`
		TableID 	string							`json:"table_id"`
		Status  	models.StaffNotificationStatus 	`json:"status"`
		CreatedAt 	time.Time 						`json:"createdAt"`
		UpdatedAt  	time.Time   					`json:"updatedAt"`
	}