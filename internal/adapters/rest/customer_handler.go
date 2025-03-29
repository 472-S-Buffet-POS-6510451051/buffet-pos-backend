package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler interface {
	Register(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	AddPoint(c *fiber.Ctx) error
	RedeemPoint(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type customerHandler struct {
	service usecases.CustomerUseCase
}

func NewCustomerHandler(service usecases.CustomerUseCase) CustomerHandler {
	return &customerHandler{
		service: service,
	}
}

// Find All Customer
// @Summary Find All Customer
// @Description Find all customers
// @Tags Loyalty
// @Accept json
// @Produce json
// @Success 200 {array} responses.BaseCustomer
// @Router /loyalty/customers [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (m *customerHandler) FindAll(c *fiber.Ctx) error {
	customers, err := m.service.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(customers)
}

// Add Points to Customer
// @Summary Add Points to Customer
// @Description Add points to a customer's account
// @Tags Loyalty
// @Accept json
// @Produce json
// @Param request body requests.CustomerAddPointRequest true "Add Point Request"
// @Success 200 {object} responses.BaseCustomer
// @Router /loyalty/add-point [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (m *customerHandler) AddPoint(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CustomerAddPointRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	
	// Add Point
	customer, err := m.service.AddPoint(c.Context(), req);
	if  err != nil {
		switch err {
		case exceptions.ErrCustomerNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not Found Customer",
			})
		case exceptions.ErrIncorrectPIN:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Incorrect PIN",
			})
		case exceptions.ErrPointLimit:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Point is full",
			})
		case exceptions.ErrInvalidPoint:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid point",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

// Redeem Points from Customer
// @Summary Redeem Points from Customer
// @Description Redeem points from a customer's account
// @Tags Loyalty
// @Accept json
// @Produce json
// @Param request body requests.CustomerRedeemRequest true "Redeem Point Request"
// @Success 200 {object} responses.BaseCustomer
// @Router /loyalty/redeem [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (m *customerHandler) RedeemPoint(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CustomerRedeemRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Redeem Point
	customer, err := m.service.RedeemPoint(c.Context(), req);
	if  err != nil {
		switch err {
		case exceptions.ErrCustomerNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not Found Customer",
			})
		case exceptions.ErrIncorrectPIN:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Incorrect PIN",
			})
		case exceptions.ErrNotEnoughPoints:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not Enough Point",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

// Register a New Customer

// @Summary Register a New Customer
// @Description Register a new customer in the system
// @Tags Loyalty
// @Accept json
// @Produce json
// @Param request body requests.CustomerRegisterRequest true "Register Request"
// @Success 200 {object} responses.SuccessResponse
// @Router /loyalty/register [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (m *customerHandler) Register(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CustomerRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Register customer
	if err := m.service.Register(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedPhone:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Phone already registered",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Customer registered successfully",
	})
}

// Delete a Customer
// @Summary Delete a Customer
// @Description Delete a customer by their ID
// @Tags Loyalty
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} responses.SuccessResponse
// @Router /loyalty/customer/{id} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (m *customerHandler) Delete(c *fiber.Ctx) error {
	id, err := utils.ValidateUUID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	if err := m.service.DeleteCustomer(c.Context(), *id); err != nil {
		switch err {
		case exceptions.ErrCustomerNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Customer not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete Customer successfully",
	})
}

