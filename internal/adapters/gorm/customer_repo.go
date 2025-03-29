package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type CustomerGormRepository struct {
	DB *gorm.DB
}

func NewCustomerGormRepository(db *gorm.DB) *CustomerGormRepository {
	return &CustomerGormRepository{
		DB: db,
	}
}

func (c *CustomerGormRepository) Create(ctx context.Context, req *requests.CustomerRegisterRequest) error {
	id, _ := uuid.NewV7()

	customer := &models.Customer{
		ID:          id,
		Phone:       req.Phone,
		PIN:         req.PIN,
	}

	result := c.DB.Create(customer)

	return result.Error
}

func (c *CustomerGormRepository) FindAll(ctx context.Context) ([]responses.BaseCustomer, error) {
	var customers []models.Customer
	if err := c.DB.Find(&customers).Error; err != nil {
		return nil, err
	}

	var response []responses.BaseCustomer
	for _, customer := range customers {
		response = append(response, responses.BaseCustomer{
			ID:    customer.ID,
			Phone: customer.Phone,
			Point:   customer.Point,
		})
	}
	return response, nil
}

func (c *CustomerGormRepository) FindByID(ctx context.Context, customerID string) (*models.Customer, error) {
	var customer models.Customer
	result := c.DB.Where("id = ?", customerID).First(&customer)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (c *CustomerGormRepository) FindByPhone(ctx context.Context, phone string) (*models.Customer, error) {
	var customer models.Customer
	result := c.DB.Where("phone = ?", phone).First(&customer)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (c *CustomerGormRepository) AddPoint(ctx context.Context, req *requests.CustomerAddPointRequest) (*responses.BaseCustomer, error) {
	var customer models.Customer
	result := c.DB.Where("phone = ?", req.Phone).First(&customer)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	customer.Point += req.Point
	if err := c.DB.Save(&customer).Error; err != nil {
		return nil, err
	}

	return &responses.BaseCustomer{
		ID:    customer.ID,
		Phone: customer.Phone,
		Point:   customer.Point,
	}, nil
}

func (c *CustomerGormRepository) RedeemPoint(ctx context.Context, req *requests.CustomerRedeemRequest, usePoint int, priceDiscount float64) (*responses.BaseCustomer, error) {
	var customer models.Customer
	result := c.DB.Where("phone = ?", req.Phone).First(&customer)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	if customer.Point < usePoint { // Check if there are enough points
		return nil, exceptions.ErrNotEnoughPoints
	}

	// ลดคะแนนของลูกค้า
	customer.Point -= usePoint
	if err := c.DB.Save(&customer).Error; err != nil {
		return nil, err
	}

	// ดึงข้อมูล Invoice ตาม InvoiceID ที่ได้รับ
	var invoice models.Invoice // สมมติว่าเรามีโมเดล Invoice
	invoiceID, err := uuid.Parse(req.InvoiceID)
	if err != nil {
		return nil, err
	}
	if err := c.DB.Where("id = ?", invoiceID).First(&invoice).Error; err != nil {
		return nil, err 
	}

	// ตั้งค่า totalPrice ใน Invoice
	totalPrice := invoice.TotalPrice
	invoice.TotalPrice = totalPrice - priceDiscount 
	if err := c.DB.Save(&invoice).Error; err != nil {
		return nil, err
	}

	return &responses.BaseCustomer{
		ID:    customer.ID,
		Phone: customer.Phone,
		Point: customer.Point,
	}, nil
}


func (c *CustomerGormRepository) Delete(ctx context.Context, customerID string) error {
	result := c.DB.Where("id = ?",customerID).Delete(&models.Customer{})
	return result.Error
}