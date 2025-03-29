package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type InvoiceHandler interface {
	GetAllUnpaidInvoices(c *fiber.Ctx) error
	GetAllPaidInvoices(c *fiber.Ctx) error
	SetInvoicePaid(c *fiber.Ctx) error
	CancelInvoice(c *fiber.Ctx) error
	CustomerGetInvoice(c *fiber.Ctx) error
	ChargeFeeFoodOverWeight(c *fiber.Ctx) error
}

type invoiceHandler struct {
	service usecases.InvoiceUseCase
}

func NewInvoiceHandler(service usecases.InvoiceUseCase) InvoiceHandler {
	return &invoiceHandler{
		service: service,
	}
}

// Get All Unpaid Invoices
// @Summary Get All Unpaid
// @Description Get all unpaid invoices.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {array} responses.InvoiceDetail
// @Router /manage/invoices/unpaid [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (i *invoiceHandler) GetAllUnpaidInvoices(c *fiber.Ctx) error {
	invoices, err := i.service.GetAllUnpaidInvoices(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(invoices)
}

// Get All Paid Invoices
// @Summary Get All Paid
// @Description Get all paid invoices.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {array} responses.InvoiceDetail
// @Router /manage/invoices/paid [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (i *invoiceHandler) GetAllPaidInvoices(c *fiber.Ctx) error {
	invoices, err := i.service.GetAllPaidInvoices(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(invoices)
}

// Set Invoice Paid
// @Summary Set Invoice Paid
// @Description Set invoice as paid.
// @Tags Manage
// @Accept json
// @Produce json
// @Param request body requests.UpdateInvoiceStatusRequest true "Update Invoice Status Request"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/invoices/set-paid [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (i *invoiceHandler) SetInvoicePaid(c *fiber.Ctx) error {
	var req *requests.UpdateInvoiceStatusRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	err := i.service.SetPaidInvoice(c.Context(), req.InvoiceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Invoice paid successfully",
	})
}

// Cancel Invoice
// @Summary Cancel Invoice
// @Description Cancel invoice.
// @Tags Manage
// @Accept json
// @Produce json
// @Param invoice_id path string true "Invoice ID"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/invoices/{invoice_id} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (i *invoiceHandler) CancelInvoice(c *fiber.Ctx) error {
	id, err := utils.ValidateUUID(c.Params("invoice_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}
	if err := i.service.DeleteInvoice(c.Context(), *id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Invoice cancelled successfully",
	})
}

// Customer Get Invoice
// @Summary Get Invoice
// @Description Get invoice by table ID.
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} responses.InvoiceDetail
// @Router /customer/invoices [get]
// @param AccessCode header string true "Access Code"
func (i *invoiceHandler) CustomerGetInvoice(c *fiber.Ctx) error {
	claims, _ := c.Locals("table").(*responses.TableDetail)
	invoice, err := i.service.CustomerGetInvoice(c.Context(), claims.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(invoice)
}

// Charge Fee Food Overweight
// @Summary Charge Fee Food Overweight
// @Description Charge Fee Food Overweight By ID
// @Tags Manage
// @Accept json
// @Produce json
// @Param request body requests.ChargeFeeFoodOverWeightRequest true "Charge Fee Food Overweight Request"
// @Success 200 {object}  responses.SuccessResponse
// @Router /manage/invoices/charge-food-overweight [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (i *invoiceHandler) ChargeFeeFoodOverWeight(c *fiber.Ctx) error {
	var req *requests.ChargeFeeFoodOverWeightRequest
	if err := c.BodyParser(&req); err != nil  {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := i.service.ChargeFeeFoodOverWeight(c.Context(), req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return  c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Charge Fee Food Overweight in invoice successfully",
	})
}

