package api

import (
	"github.com/gofiber/fiber/v2"
	"royalty-api/internal/domain"
	"royalty-api/internal/service"
	"strconv"
)

type userHandler struct {
	Service service.UserService
}

func NewUserHandler(serv service.UserService) userHandler {
	return userHandler{Service: serv}
}

// listUsers function
// @Summary Get list of users
// @Description Retrieve list of users
// @Tags User
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /user [get]
func (u *userHandler) listUsers(c *fiber.Ctx) error {
	results, err := u.Service.List()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Get List Users Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}

// viewUser function
// @Summary Get user detail
// @Description Retrieve detail of users
// @Tags User
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /user/:id [get]
func (u *userHandler) viewUser(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, _ := strconv.Atoi(idStr)

	result, err := u.Service.View(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Get List Movie Success",
		Data:       result,
		StatusCode: fiber.StatusOK,
	})
}

// insertUser function
// @Summary Register User
// @Description Register User to the App
// @Tags User
// @Accept json
// @Produce json
// @Param body body domain.UserRequest true "User information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /user [post]
func (u *userHandler) insertUser(c *fiber.Ctx) error {
	var bodyRequest domain.UserRequest

	if err := c.BodyParser(&bodyRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	results, err := u.Service.Insert(&bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Insert User Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}

// updateUser function
// @Summary Update User
// @Description Update Exist User in the Application
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body domain.UserRequest true "User information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /user/:id [patch]
func (u *userHandler) updateUser(c *fiber.Ctx) error {
	var bodyRequest domain.UserRequest

	idStr := c.Params("id")

	if err := c.BodyParser(&bodyRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	id, _ := strconv.Atoi(idStr)

	bodyRequest.ID = uint(id)

	result, err := u.Service.Update(&bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Update User Success",
		Data:       result,
		StatusCode: fiber.StatusOK,
	})
}

// viewUser function
// @Summary Delete user
// @Description Delete user from database
// @Tags User
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /user/:id [delete]
func (u *userHandler) deleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	result, err := u.Service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Delete User Success",
		Data:       result,
		StatusCode: fiber.StatusOK,
	})
}
