package model

import (
	"gorm.io/gorm"
	"time"
)

// Transaction TODO: Add detail transaction where contain products
type Transaction struct {
	ID               uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	InvoiceNumber    string    `json:"invoice_number" gorm:"type:varchar"`
	TransactionDate  time.Time `json:"transaction_date"`
	TotalAmount      int64     `json:"total_amount" gorm:"type:bigint"`
	TenantID         uint      `json:"-"`
	Tenant           Tenant    `json:"tenant" gorm:"foreignKey:TenantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID           uint      `json:"-"`
	User             User      `json:"user" gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VoucherGenerated bool      `json:"voucher_generated"`
	Deleted          bool      `json:"deleted" gorm:"type:boolean; default:false"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (u *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}
