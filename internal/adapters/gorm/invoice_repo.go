package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceGormRepository struct {
	DB *gorm.DB
}

func NewInvoiceGormRepository(db *gorm.DB) *InvoiceGormRepository {
	return &InvoiceGormRepository{
		DB: db,
	}
}

func (i *InvoiceGormRepository) FindByID(ctx context.Context, invoiceID string) (*models.Invoice, error) {
	invoiceIDParse, err := uuid.Parse(invoiceID)
	if err != nil {
		return nil, err
	}
	invoice := &models.Invoice{}
	result := i.DB.First(invoice, invoiceIDParse)

	return invoice, result.Error
}

func (i *InvoiceGormRepository) Create(ctx context.Context, tableID string, totalPrice float64, peopleAmount int) error {
	id, _ := uuid.NewV7()

	tableIDParse, err := uuid.Parse(tableID)
	if err != nil {
		return err
	}

	invoice := &models.Invoice{
		ID:           id,
		TableID:      &tableIDParse,
		PeopleAmount: peopleAmount,
		TotalPrice:   totalPrice,
		IsPaid:       false,
	}
	result := i.DB.Create(invoice)
	return result.Error
}

func (i *InvoiceGormRepository) SetPaid(ctx context.Context, invoiceID string) error {
	invoiceIDParse, err := uuid.Parse(invoiceID)
	if err != nil {
		return err
	}
	result := i.DB.Model(&models.Invoice{}).Where("id = ?", invoiceIDParse).Update("is_paid", true)
	return result.Error
}

func (i *InvoiceGormRepository) Cancel(ctx context.Context, invoiceID string) error {
	invoiceIDParse, err := uuid.Parse(invoiceID)
	if err != nil {
		return err
	}

	result := i.DB.Delete(&models.Invoice{}, invoiceIDParse)
	return result.Error
}

func (i *InvoiceGormRepository) GetAllUnpaid(ctx context.Context) ([]models.Invoice, error) {
	var invoices []models.Invoice
	result := i.DB.Where("is_paid = ?", false).Order("created_at desc").Find(&invoices)
	return invoices, result.Error
}

func (i *InvoiceGormRepository) GetAllPaid(ctx context.Context) ([]models.Invoice, error) {
	var invoices []models.Invoice
	result := i.DB.Where("is_paid = ?", true).Order("created_at desc").Find(&invoices)
	return invoices, result.Error
}

func (i *InvoiceGormRepository) GetByTableID(ctx context.Context, tableID string) (*models.Invoice, error) {
	tableIDParse, err := uuid.Parse(tableID)
	if err != nil {
		return nil, err
	}
	invoice := &models.Invoice{}
	result := i.DB.Where("table_id = ?", tableIDParse).Order("created_at desc").First(invoice)
	return invoice, result.Error
}

func (i *InvoiceGormRepository) AddTotalPriceByID(ctx context.Context, invoiceID string, price float64) error {
	var invoice models.Invoice
	result := i.DB.Where("id = ?", invoiceID).First(&invoice)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	
	invoice.TotalPrice += price

	if err := i.DB.Save(&invoice).Error; err != nil {
		return err
	}

	return nil
}

func (i *InvoiceGormRepository) SetPriceFeeFoodOverWeightByID(ctx context.Context, invoiceID string, price float64) error {
	var invoice models.Invoice
	result := i.DB.Where("id = ?", invoiceID).First(&invoice)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	
	invoice.PriceFeeFoodOverWeight = price

	if err := i.DB.Save(&invoice).Error; err != nil {
		return err
	}

	return nil
}

func (i *InvoiceGormRepository) SetCustomerPhoneByID(ctx context.Context, invoiceID string, customerPhone string) error {
	var invoice models.Invoice
	result := i.DB.Where("id = ?", invoiceID).First(&invoice)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	
	invoice.CustomerPhone = customerPhone

	if err := i.DB.Save(&invoice).Error; err != nil {
		return err
	}

	return nil
}
