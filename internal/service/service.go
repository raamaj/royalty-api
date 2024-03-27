package service

import "royalty-api/internal/repository"

type Services struct {
	UserService        UserService
	VoucherService     VoucherService
	TransactionService TransactionService
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		UserService:        NewUserService(repo.User),
		VoucherService:     NewVoucherService(repo.Voucher),
		TransactionService: NewTransactionService(repo.Transaction),
	}
}
