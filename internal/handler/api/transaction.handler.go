package api

import (
	"github.com/gofiber/fiber/v2"
	"royalty-api/internal/domain"
	"royalty-api/internal/service"
	"strconv"
	"time"
)

type transactionHandler struct {
	Service service.TransactionService
}

func NewTransactionHandler(serv service.TransactionService) transactionHandler {
	return transactionHandler{Service: serv}
}

// createTransaction function
// @Summary Create Transaction for each User
// @Description Create Transaction for each User
// @Tags Transaction
// @Accept json
// @Produce json
// @Param body body domain.TransactionRequest true "Transaction information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /transaction [post]
func (t *transactionHandler) createTransaction(c *fiber.Ctx) error {
	var body domain.TransactionRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	result, err := t.Service.Create(&body, time.Now())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Create Transaction Success",
		Data:       result,
		StatusCode: fiber.StatusOK,
	})
}

// listTransaction function
// @Summary List of Transaction for each User
// @Description List of Transaction for each User
// @Tags Transaction
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /transaction/:userId [get]
func (t *transactionHandler) listTransaction(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userID, _ := strconv.Atoi(userIdStr)

	results, err := t.Service.List(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Create Transaction Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}
