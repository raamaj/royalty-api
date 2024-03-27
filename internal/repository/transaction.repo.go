package repository

import (
	"errors"
	"gorm.io/gorm"
	"royalty-api/internal/domain"
	"royalty-api/internal/model"
	"time"
)

type transactionRepository struct {
	db *gorm.DB
}

func (t transactionRepository) UpdateVoucherGenerated(invoice string) (*model.Transaction, error) {
	var transaction model.Transaction

	db := t.db.Model(&transaction)

	checkTransaction := db.Debug().First(&transaction, "InvoiceNumber=?", invoice)
	if checkTransaction.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	if transaction.ID < 1 {
		return nil, errors.New("Unknown transaction with Invoice Number " + invoice)
	}

	transaction.VoucherGenerated = true
	transaction.UpdatedAt = time.Now()

	updatedTransaction := db.Debug().Updates(&transaction)
	if updatedTransaction.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &transaction, nil
}

func (t transactionRepository) Create(request *domain.TransactionRequest, timeNow time.Time) (*model.Transaction, error) {
	var transaction model.Transaction
	var tenant model.Tenant
	var user model.User

	tenant.ID = request.TenantID

	db := t.db.Model(&tenant)

	checkTenant := db.Debug().First(&tenant, "deleted=false")
	if checkTenant.Error != nil {
		return nil, checkTenant.Error
	}

	user.ID = request.UserID

	db = t.db.Model(&user)

	checkUser := db.Debug().First(&user, "deleted=false")
	if checkUser.Error != nil {
		return nil, checkUser.Error
	}

	transaction.TransactionDate = request.TransactionDate
	transaction.InvoiceNumber = request.InvoiceNumber
	transaction.Tenant = tenant
	transaction.User = user
	transaction.CreatedAt = timeNow
	transaction.TotalAmount = request.TotalAmount
	transaction.VoucherGenerated = request.VoucherGenerated
	transaction.Deleted = false

	db = t.db.Model(&transaction)

	addTransaction := db.Debug().Create(&transaction)
	if addTransaction.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &transaction, nil
}

func (t transactionRepository) List(userId uint) (*[]model.Transaction, error) {
	var transactions []model.Transaction

	db := t.db.Model(&transactions)

	checkTransactions := db.Debug().Find(&transactions, "user_id=?", userId)
	if checkTransactions.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &transactions, nil
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}
