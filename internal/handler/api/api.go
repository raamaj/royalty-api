package api

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	_ "royalty-api/docs"
	"royalty-api/internal/repository"
	"royalty-api/internal/service"
)

func NewApi(app *fiber.App, db *gorm.DB) {
	app.Use(recover2.New())
	app.Use(requestid.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	repos := repository.NewRepository(db)
	servs := service.NewServices(repos)

	userHandl := NewUserHandler(servs.UserService)
	voucherHandl := NewVoucherHandler(servs.VoucherService)
	transactionHandl := NewTransactionHandler(servs.TransactionService)

	api := app.Group("/api")

	// User API
	user := api.Group("/user")
	user.Get("/", userHandl.listUsers)
	user.Get("/:id", userHandl.viewUser)
	user.Post("/", userHandl.insertUser)
	user.Patch("/:id", userHandl.updateUser)
	user.Delete("/:id", userHandl.deleteUser)

	// Voucher API
	voucher := api.Group("/voucher")
	voucher.Get("/generate/:invoice", voucherHandl.createVoucher)
	voucher.Get("/redeem/:code", voucherHandl.redeemVoucher)
	voucher.Get("/:userId", voucherHandl.listVouchers)

	// Transaction API
	transaction := api.Group("/transaction")
	transaction.Post("/", transactionHandl.createTransaction)
	transaction.Get("/:userId", transactionHandl.listTransaction)

	log.Fatal(app.Listen(viper.GetString("APP_ADDRESS")))
}
