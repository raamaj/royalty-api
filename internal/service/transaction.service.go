package service

import (
	"math/rand"
	"royalty-api/internal/domain"
	"royalty-api/internal/model"
	"royalty-api/internal/repository"
	"time"
)

type TransactionServiceImpl struct {
	Repository repository.TransactionRepository
}

func (t TransactionServiceImpl) Create(request *domain.TransactionRequest, timeNow time.Time) (*domain.TransactionResponse, error) {

	request.VoucherGenerated = false
	request.InvoiceNumber = "INV-" + t.generateInvoice()

	transaction, err := t.Repository.Create(request, timeNow)
	if err != nil {
		return nil, err
	}

	return t.convertToResponse(transaction), err
}

func (t TransactionServiceImpl) List(userId uint) (*[]domain.TransactionResponse, error) {
	transactions, err := t.Repository.List(userId)
	if err != nil {
		return nil, err
	}

	return t.convertToResponses(transactions), nil
}

func (t TransactionServiceImpl) convertToResponse(trans *model.Transaction) *domain.TransactionResponse {
	return &domain.TransactionResponse{
		ID:               trans.ID,
		InvoiceNumber:    trans.InvoiceNumber,
		TransactionDate:  trans.TransactionDate,
		TotalAmount:      trans.TotalAmount,
		TenantID:         trans.TenantID,
		UserID:           trans.UserID,
		GeneratedVoucher: trans.VoucherGenerated,
		CreatedAt:        trans.CreatedAt,
		UpdatedAt:        trans.UpdatedAt,
	}
}

func (t TransactionServiceImpl) convertToResponses(transactions *[]model.Transaction) *[]domain.TransactionResponse {
	var response []domain.TransactionResponse
	for _, trans := range *transactions {
		response = append(response, domain.TransactionResponse{
			ID:               trans.ID,
			InvoiceNumber:    trans.InvoiceNumber,
			TransactionDate:  trans.TransactionDate,
			TotalAmount:      trans.TotalAmount,
			TenantID:         trans.TenantID,
			UserID:           trans.UserID,
			GeneratedVoucher: trans.VoucherGenerated,
			CreatedAt:        trans.CreatedAt,
			UpdatedAt:        trans.UpdatedAt,
		})
	}

	return &response
}

func (t TransactionServiceImpl) generateInvoice() string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewTransactionService(transactionRepo repository.TransactionRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		Repository: transactionRepo,
	}
}
