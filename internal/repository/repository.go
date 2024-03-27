package repository

import "gorm.io/gorm"

type Repositories struct {
	User        UserRepository
	Voucher     VoucherRepository
	Transaction TransactionRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:        NewUserRepository(db),
		Voucher:     NewVoucherRepository(db),
		Transaction: NewTransactionRepository(db),
	}
}
