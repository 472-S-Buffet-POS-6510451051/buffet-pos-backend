package main

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/gorm"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/middleware"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/rest"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/infrastructure/cloudinary"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "github.com/cs471-buffetpos/buffet-pos-backend/docs"
)

// @title BuffetPOS API
// @description This is the BuffetPOS API documentation.
// @version 1.0
func main() {
	app := fiber.New()
	cfg := configs.NewConfig()

	db := bootstrap.NewDB(cfg)

	cld := cloudinary.NewCloudinaryStorageService(cfg)

	settingRepo := gorm.NewSettingGormRepository(db)
	settingService := usecases.NewSettingService(settingRepo, cfg)
	settingHandler := rest.NewSettingHandler(settingService)

	userRepo := gorm.NewUserGormRepository(db)
	invoiceRepo := gorm.NewInvoiceGormRepository(db)
	tableRepo := gorm.NewTableGormRepository(db)
	categoryRepo := gorm.NewCategoryGormRepository(db)
	menuRepo := gorm.NewMenuGormRepository(db)
	orderRepo := gorm.NewOrderGormRepository(db)
	orderItemRepo := gorm.NewOrderItemGormRepository(db)
	customerRepo := gorm.NewCustomerGormRepository(db)
	staffNotiRepo := gorm.NewStaffNotificationGormRepository(db)

	userService := usecases.NewUserService(userRepo, cfg)
	userHandler := rest.NewUserHandler(userService)

	invoiceService := usecases.NewInvoiceService(invoiceRepo, tableRepo, orderRepo, settingRepo, cfg)
	tableService := usecases.NewTableService(tableRepo, invoiceRepo, settingRepo, cfg)

	tableHandler := rest.NewTableHandler(tableService)
	invoiceHandler := rest.NewInvoiceHandler(invoiceService)

	categoryService := usecases.NewCategoryService(categoryRepo, cfg)
	categoryHandler := rest.NewCategoryHandler(categoryService)

	menuService := usecases.NewMenuService(menuRepo, cfg, cld)
	menuHandler := rest.NewMenuHandler(menuService)

	orderService := usecases.NewOrderService(orderRepo, orderItemRepo, menuRepo, cfg)
	orderHandler := rest.NewOrderHandler(orderService)

	customerService := usecases.NewCustomerService(customerRepo, invoiceRepo, settingRepo, cfg)
	customerHandler := rest.NewCustomerHandler(customerService)

	orderItemService := usecases.NewOrderItemService(orderItemRepo, cfg)
	orderItemHandler := rest.NewOrderItemHandler(orderItemService)

	staffNotiService := usecases.NewStaffNotificationService(staffNotiRepo, cfg)
	staffNotiHandler := rest.NewStaffNotificationHandler(staffNotiService)

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BuffetPOS is running 🎉")
	})

	auth := app.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	general := app.Group("/general")
	general.Get("/menus/best-selling", orderItemHandler.GetBestSellingMenu)

	customer := app.Group("/customer", middleware.CustomerMiddleware(cfg, tableService))
	customer.Get("/menus", menuHandler.CustomerFindAll)
	customer.Get("/menus/:id", menuHandler.CustomerFindByID)
	customer.Post("/orders", orderHandler.CreateOrder)
	customer.Get("/orders/history", orderHandler.CustomerGetOrderHistory)
	customer.Get("/tables", tableHandler.FindCustomerTable)
	customer.Get("/categories", categoryHandler.FindAllCategories)
	customer.Get("/invoices", invoiceHandler.CustomerGetInvoice)
	// staff-notification : calling staff
	customer.Post("/staff-notifications", staffNotiHandler.NotifyStaff)
	customer.Get("/staff-notifications/:table_id", staffNotiHandler.GetStaffNotificationByTableId)

	loyalty := app.Group("/loyalty", middleware.AuthMiddleware(cfg), middleware.RoleMiddleware(models.Employee, models.Manager))
	loyalty.Post("/register", customerHandler.Register)
	loyalty.Get("/customers", customerHandler.FindAll)
	loyalty.Post("/add-point", customerHandler.AddPoint)
	loyalty.Post("/redeem", customerHandler.RedeemPoint)
	loyalty.Delete("/customer/:id", customerHandler.Delete)

	manage := app.Group("/manage", middleware.AuthMiddleware(cfg), middleware.RoleMiddleware(models.Employee, models.Manager))
	manage.Get("/tables", tableHandler.FindAllTables)
	manage.Get("/tables/:id", tableHandler.FindTableByID)
	manage.Post("/tables", tableHandler.AddTable)
	manage.Put("/tables", tableHandler.Edit)
	manage.Delete("/tables/:id", tableHandler.Delete)
	manage.Post("/tables/assign", tableHandler.AssignTable)

	manage.Get("/invoices/paid", invoiceHandler.GetAllPaidInvoices)
	manage.Get("/invoices/unpaid", invoiceHandler.GetAllUnpaidInvoices)
	manage.Put("/invoices/set-paid", invoiceHandler.SetInvoicePaid)
	manage.Delete("/invoices/:invoice_id", invoiceHandler.CancelInvoice)
	manage.Put("/invoices/charge-food-overweight", invoiceHandler.ChargeFeeFoodOverWeight)

	manage.Get("/categories", categoryHandler.FindAllCategories)
	manage.Get("/categories/:id", categoryHandler.FindCategoryByID)
	manage.Post("/categories", categoryHandler.AddCategory)
	manage.Delete("/categories/:id", categoryHandler.DeleteCategory)

	manage.Get("/menus", menuHandler.FindAll)

	manage.Get("/menus/:id", menuHandler.FindByID)
	manage.Post("/menus", menuHandler.Create)
	manage.Put("/menus", menuHandler.Edit)
	manage.Delete("/menus/:id", menuHandler.Delete)

	manage.Get("/orders", orderHandler.GetOrdersByStatus)
	manage.Get("/orders/tables", orderHandler.GetOrdersByTableID)
	manage.Put("/orders/status", orderHandler.UpdateOrderStatus)

	manage.Get("/settings/price-per-person", settingHandler.GetPricePerPerson)
	manage.Put("/settings/price-per-person", settingHandler.SetPricePerPerson)
	manage.Get("/settings/use-point-per-person", settingHandler.GetUsePointPerPerson)
	manage.Put("/settings/use-point-per-person", settingHandler.SetUsePointPerPerson)
	manage.Get("/settings/price-fee-food-overweight", settingHandler.GetPriceFeeFoodOverWeight)
	manage.Put("/settings/price-fee-food-overweight", settingHandler.SetPriceFeeFoodOverWeight)

	// staff notification
	manage.Get("/staff-notifications", staffNotiHandler.GetAllStaffNotification)
	manage.Get("/staff-notifications/:status", staffNotiHandler.GetAllStaffNotificationByStatus)
	manage.Put("/staff-notifications", staffNotiHandler.UpdateStaffNotificationStatus)

	app.Listen(":3001")
}
