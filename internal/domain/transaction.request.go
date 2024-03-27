package domain

import "time"

type TransactionRequest struct {
	ID               uint      `json:"id"`
	InvoiceNumber    string    `json:"invoice_number"`
	TransactionDate  time.Time `json:"transaction_date"`
	TotalAmount      int64     `json:"total_amount"`
	TenantID         uint      `json:"tenant_id"`
	UserID           uint      `json:"user_id"`
	VoucherGenerated bool
}

type TransactionResponse struct {
	ID               uint      `json:"id"`
	InvoiceNumber    string    `json:"invoice_number"`
	TransactionDate  time.Time `json:"transaction_date"`
	TotalAmount      int64     `json:"total_amount"`
	TenantID         uint      `json:"tenant_id"`
	UserID           uint      `json:"user_id"`
	GeneratedVoucher bool      `json:"generated_voucher"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
