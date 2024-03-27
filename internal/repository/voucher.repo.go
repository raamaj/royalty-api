package repository

import (
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"royalty-api/internal/model"
	"time"
)

type voucherRepository struct {
	db *gorm.DB
}

func (v voucherRepository) List(userID uint) (*[]model.Voucher, error) {
	var vouchers []model.Voucher

	db := v.db.Model(&vouchers)

	checkVouchers := db.Debug().Find(&vouchers, "UserID=?", userID)
	if checkVouchers.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &vouchers, nil
}

func (v voucherRepository) Create(invoice string, timeNow time.Time) (*model.Voucher, error) {
	var transaction model.Transaction
	var voucher model.Voucher
	tx := v.db.Begin()

	checkTrans := tx.Debug().Find(&transaction, "invoice_number=?", invoice)
	if checkTrans.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	if transaction.ID < 1 {
		return nil, errors.New("Unknown Transaction with Invoice Number " + invoice + "!")
	}

	if transaction.TotalAmount < 1000000 {
		return nil, errors.New("Cannot generate voucher for Invoice Number " + invoice + " because Total Amount is lower than Rp. 1.000.000")
	}

	voucher.Code = v.generateCode()

	checkVoucher := tx.First(&voucher, "code=?", voucher.Code)
	if checkVoucher.RowsAffected > 1 {
		return nil, errors.New("Voucher with code " + voucher.Code + " is already registered.")
	}

	voucher.UserID = transaction.UserID
	voucher.Used = false
	voucher.Value = 10000
	voucher.ExpiredTime = timeNow.AddDate(0, 3, 0)
	voucher.CreatedAt = timeNow
	voucher.Deleted = false

	checkVoucher = tx.Debug().Create(&voucher)
	if checkVoucher.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	transaction.VoucherGenerated = true
	transaction.UpdatedAt = timeNow

	updatedTrans := tx.Debug().Updates(&transaction)
	if updatedTrans.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	tx.Commit()

	return &voucher, nil
}

func (v voucherRepository) Use(code string, timeNow time.Time) (*model.Voucher, error) {
	var voucher model.Voucher

	db := v.db.Model(&voucher)

	voucher.Code = code

	existVoucher := db.Debug().First(&voucher, "code=? AND deleted=false", code)
	if existVoucher.Error != nil {
		return nil, existVoucher.Error
	}

	if voucher.Used {
		return nil, errors.New("Voucher with code " + code + " has been used by other")
	}

	if timeNow.After(voucher.ExpiredTime) {
		return nil, errors.New("Voucher with code " + code + " has been expired")
	}

	voucher.Used = true
	voucher.UpdatedAt = timeNow

	updatedVoucher := db.Debug().Model(&voucher).Updates(&voucher)
	if updatedVoucher.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service")
	}

	return &voucher, nil
}

func (v voucherRepository) generateCode() string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewVoucherRepository(db *gorm.DB) *voucherRepository {
	return &voucherRepository{db: db}
}
