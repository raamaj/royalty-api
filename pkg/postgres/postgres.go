package postgres

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"royalty-api/internal/model"
	"royalty-api/pkg/logger"
	"time"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed error connect to database")
	}

	err = db.AutoMigrate(&model.User{}, &model.Transaction{}, &model.Voucher{})
	seed(db)

	if err != nil {
		defer logger.Info("Database connection failed")
		logger.Error("error + ", zap.Error(err))

		return nil, fmt.Errorf("faield connected to database")
	}

	return db, nil
}

func seed(gorm *gorm.DB) {
	var tenants []model.Tenant
	db := gorm.Model(&tenants)

	err := db.Find(&tenants)
	if err.Error != nil {
		logger.Error("Seed Error : Get Tetants")
	}

	if len(tenants) < 1 {
		tenant := model.Tenant{
			Name:      "Matahari",
			Deleted:   false,
			UpdatedAt: time.Time{},
		}

		err = db.Create(&tenant)
		if err.Error != nil {
			logger.Error("Seed Error : Insert Tenant")
		}
	}
}
