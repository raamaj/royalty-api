package repository

import (
	"royalty-api/internal/domain"
	"royalty-api/internal/model"
	"time"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type UserRepository interface {
	List() (*[]model.User, error)
	View(id uint) (*model.User, error)
	Insert(input *domain.UserRequest, timeNow time.Time) (*model.User, error)
	Update(input *domain.UserRequest, timeNow time.Time) (*model.User, error)
	Delete(id uint, timeNow time.Time) (*model.User, error)
}

type VoucherRepository interface {
	Create(invoice string, timeNow time.Time) (*model.Voucher, error)
	Use(code string, timeNow time.Time) (*model.Voucher, error)
	List(userID uint) (*[]model.Voucher, error)
}

type TransactionRepository interface {
	Create(request *domain.TransactionRequest, timeNow time.Time) (*model.Transaction, error)
	List(userId uint) (*[]model.Transaction, error)
	UpdateVoucherGenerated(invoice string) (*model.Transaction, error)
}
