package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type StaffNotificationHandler interface {
	NotifyStaff(c *fiber.Ctx) error
	GetAllStaffNotification(c *fiber.Ctx) error
	GetAllStaffNotificationByStatus(c *fiber.Ctx) error
	UpdateStaffNotificationStatus(c *fiber.Ctx) error
	GetStaffNotificationByTableId(c *fiber.Ctx) error
}

type staffNotificationHandler struct {
	service usecases.StaffNotificationUseCase
}

func NewStaffNotificationHandler(service usecases.StaffNotificationUseCase) StaffNotificationHandler {
	return &staffNotificationHandler{
		service: service,
	}
}

// Get All Staff Notification
// @Summary Get All Staff Notification
// @Description Get All Staff Notification.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {array} responses.StaffNotificationBase
// @Router /manage/staff-notifications [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (s *staffNotificationHandler) GetAllStaffNotification(c *fiber.Ctx) error {
	notifications, err := s.service.GetAllStaffNotification(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(notifications)
}

// Get All Staff Notification By Status
// @Summary Get All Staff Notification By Status
// @Description Get All Staff Notification By Status
// @Tags  Manage
// @Accept json
// @Produce json
// @Param status path string true "Status Notification"
// @Success 200 {object} responses.StaffNotificationBase
// @Router /manage/staff-notifications/{status} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (s *staffNotificationHandler) GetAllStaffNotificationByStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	notifications, err := s.service.GetAllStaffNotificationByStatus(c.Context(), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(notifications)
}

// Notify Staff
// @Summary Notify to Staff
// @Description Notify to Staff
// @Tags Customer
// @Accept json
// @Param request body requests.StaffNotificationRequest true "Notify to Staff request"
// @Success 201 {object} responses.SuccessResponse
// @Router /customer/staff-notifications [post]
// @param AccessCode header string true "Access Code"
func (s *staffNotificationHandler) NotifyStaff(c *fiber.Ctx) error {
	var req requests.StaffNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := s.service.NotifyStaff(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Notification sent successfully",
	})
}

// Update Status
// @Summary Update Status Staff Notification
// @Description Set invoice as paid.
// @Tags Manage
// @Accept json
// @Param request body requests.UpdateStaffNotificationRequest true "Update Staff Notification Request"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/staff-notifications [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (s *staffNotificationHandler) UpdateStaffNotificationStatus(c *fiber.Ctx) error {
	var req requests.UpdateStaffNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := s.service.UpdateStatus(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Status updated successfully",
	})
}

// Get Staff Notification By Table ID
// @Summary Get Staff Notification By Table ID
// @Description Get Staff Notification By Table ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param table_id path string true "Table ID"
// @Success 200 {object} responses.StaffNotificationBase
// @Router /customer/staff-notifications/{table_id} [get]
// @param AccessCode header string true "Access Code"
func (s *staffNotificationHandler) GetStaffNotificationByTableId(c *fiber.Ctx) error {
	tableID := c.Params("table_id")
	notification, err := s.service.GetStaffNotificationByTableId(c.Context(), tableID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(notification)
}
