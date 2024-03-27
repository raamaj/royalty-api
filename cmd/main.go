package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"royalty-api/internal/handler/api"
	"royalty-api/pkg/dotenv"
	"royalty-api/pkg/logger"
	"royalty-api/pkg/postgres"
)

// @title Royalty API
// @version 1.0
// @description RESTFul API to distribute and redeem the voucher
// @contact.name Rama Jayapermana
// @contact.email jayapermanarama@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	err := dotenv.Viper()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	db, err := postgres.NewClient()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	app := fiber.New()
	api.NewApi(app, db)
}
