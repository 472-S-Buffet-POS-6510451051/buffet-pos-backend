package main

import (
	"log"

	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.NewConfig()
	db := bootstrap.NewDB(cfg)

	if err := db.AutoMigrate(
		&models.User{},
		&models.Table{},
		&models.Category{},
		&models.Menu{},
		&models.Setting{},
		&models.Invoice{},
		&models.Order{},
		&models.OrderItem{},
		&models.Customer{},
		&models.StaffNotification{},
	); err != nil {
		log.Fatal(err)
	}

	// Seed default settings
	if err := initializeDefaultSettings(db); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Migration completed")
}

func initializeDefaultSettings(db *gorm.DB) error {
	settings := []models.Setting{
		{
			Key: "pricePerPerson", 
			Value: "250.00",
		},
		{
			Key: "usePointPerPerson",
			Value: "10",
		},
		{
			Key: "priceFeeFoodOverWeight",
			Value : "10",
		},	
	}

	for _, setting := range settings {
		var existingSetting models.Setting
		if err := db.Where("key = ?", setting.Key).First(&existingSetting).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&setting).Error; err != nil {
					return err
				}
				log.Printf("✅ Inserted default setting for %s: %s\n", setting.Key, setting.Value)
			} else {
				return err
			}
		}
	}
	return nil
}
