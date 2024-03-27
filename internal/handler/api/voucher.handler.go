package api

import (
	"github.com/gofiber/fiber/v2"
	"royalty-api/internal/domain"
	"royalty-api/internal/service"
	"strconv"
	"time"
)

type voucherHandler struct {
	Service service.VoucherService
}

func NewVoucherHandler(serv service.VoucherService) voucherHandler {
	return voucherHandler{Service: serv}
}

// createVoucher function
// @Summary Generate Voucher for User
// @Description Generate Voucher for User based on Invoice
// @Tags Voucher
// @Produce json
// @Param invoice path string true "Invoice Number"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /voucher/generate/:invoice [get]
func (v *voucherHandler) createVoucher(c *fiber.Ctx) error {
	invoice := c.Params("invoice")
	result, err := v.Service.Create(invoice, time.Now())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Create Voucher Success",
		Data:       result,
		StatusCode: fiber.StatusOK,
	})
}

// redeemVoucher function
// @Summary Redeem Voucher for User
// @Description Redeem Voucher for User
// @Tags Voucher
// @Produce json
// @Param code path string true "Voucher Code"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /voucher/redeem/:code [get]
func (v *voucherHandler) redeemVoucher(c *fiber.Ctx) error {
	code := c.Params("code")
	result, err := v.Service.Redeem(code, time.Now())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Redeem Voucher Success",
		Data:       result,
		StatusCode: fiber.StatusOK,
	})
}

// listVouchers function
// @Summary Get list of vouchers
// @Description Retrieve list of vouchers based on User ID
// @Tags Voucher
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /voucher/:userId [get]
func (v *voucherHandler) listVouchers(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, _ := strconv.Atoi(userIdStr)
	results, err := v.Service.List(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Get List Vouchers Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}
