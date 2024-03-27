package service

import (
	"royalty-api/internal/domain"
	"time"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type UserService interface {
	List() (*[]domain.UserResponse, error)
	View(id uint) (*domain.UserResponse, error)
	Insert(input *domain.UserRequest) (*domain.UserResponse, error)
	Update(input *domain.UserRequest) (*domain.UserResponse, error)
	Delete(id uint) (*domain.UserResponse, error)
}

type VoucherService interface {
	Create(invoice string, timeNow time.Time) (*domain.VoucherResponse, error)
	Redeem(code string, timeNow time.Time) (*domain.VoucherResponse, error)
	List(userID uint) (*[]domain.VoucherResponse, error)
}

type TransactionService interface {
	Create(request *domain.TransactionRequest, timeNow time.Time) (*domain.TransactionResponse, error)
	List(userId uint) (*[]domain.TransactionResponse, error)
}
