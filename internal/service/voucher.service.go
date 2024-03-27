package service

import (
	"royalty-api/internal/domain"
	"royalty-api/internal/model"
	"royalty-api/internal/repository"
	"time"
)

type VoucherServiceImpl struct {
	Repository repository.VoucherRepository
}

func (v VoucherServiceImpl) Create(invoice string, timeNow time.Time) (*domain.VoucherResponse, error) {
	voucher, err := v.Repository.Create(invoice, timeNow)
	if err != nil {
		return nil, err
	}

	return v.convertToResponse(voucher), err
}

func (v VoucherServiceImpl) Redeem(code string, timeNow time.Time) (*domain.VoucherResponse, error) {
	voucher, err := v.Repository.Use(code, timeNow)
	if err != nil {
		return nil, err
	}

	return v.convertToResponse(voucher), err
}

func (v VoucherServiceImpl) List(userID uint) (*[]domain.VoucherResponse, error) {
	vouchers, err := v.Repository.List(userID)
	if err != nil {
		return nil, err
	}

	return v.convertToResponses(vouchers), err
}

func (v VoucherServiceImpl) convertToResponse(voucher *model.Voucher) *domain.VoucherResponse {
	return &domain.VoucherResponse{
		ID: voucher.ID,
		User: domain.UserResponse{
			ID:        voucher.User.ID,
			Name:      voucher.User.Name,
			Phone:     voucher.User.Phone,
			Address:   voucher.User.Address,
			Email:     voucher.User.Email,
			Username:  voucher.User.Username,
			CreatedAt: voucher.User.CreatedAt,
			UpdatedAt: voucher.User.UpdatedAt,
		},
		Name:        voucher.Name,
		Code:        voucher.Code,
		Value:       voucher.Value,
		Used:        voucher.Used,
		ExpiredTime: voucher.ExpiredTime,
		CreatedAt:   voucher.CreatedAt,
		UpdatedAt:   voucher.UpdatedAt,
	}
}

func (v VoucherServiceImpl) convertToResponses(vouchers *[]model.Voucher) *[]domain.VoucherResponse {
	var responses []domain.VoucherResponse
	for _, voucher := range *vouchers {
		responses = append(responses, domain.VoucherResponse{
			ID: voucher.ID,
			User: domain.UserResponse{
				ID:        voucher.User.ID,
				Name:      voucher.User.Name,
				Phone:     voucher.User.Phone,
				Address:   voucher.User.Address,
				Email:     voucher.User.Email,
				Username:  voucher.User.Username,
				CreatedAt: voucher.User.CreatedAt,
				UpdatedAt: voucher.User.UpdatedAt,
			},
			Name:        voucher.Name,
			Code:        voucher.Code,
			Value:       voucher.Value,
			Used:        voucher.Used,
			ExpiredTime: voucher.ExpiredTime,
			CreatedAt:   voucher.CreatedAt,
			UpdatedAt:   voucher.UpdatedAt,
		})
	}

	return &responses
}

func NewVoucherService(voucherRepo repository.VoucherRepository) *VoucherServiceImpl {
	return &VoucherServiceImpl{
		Repository: voucherRepo,
	}
}
