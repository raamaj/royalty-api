package service

import (
	"royalty-api/internal/domain"
	"royalty-api/internal/model"
	"royalty-api/internal/repository"
	"time"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func (u UserServiceImpl) List() (*[]domain.UserResponse, error) {
	users, err := u.Repository.List()
	if err != nil {
		return nil, err
	}

	return u.convertToResponses(users), err
}

func (u UserServiceImpl) View(id uint) (*domain.UserResponse, error) {
	user, err := u.Repository.View(id)
	if err != nil {
		return nil, err
	}

	return u.convertToResponse(user), err
}

func (u UserServiceImpl) Insert(input *domain.UserRequest) (*domain.UserResponse, error) {
	createdUser, err := u.Repository.Insert(input, time.Now())
	if err != nil {
		return nil, err
	}

	return u.convertToResponse(createdUser), err
}

func (u UserServiceImpl) Update(input *domain.UserRequest) (*domain.UserResponse, error) {
	updatedUser, err := u.Repository.Update(input, time.Now())
	if err != nil {
		return nil, err
	}

	return u.convertToResponse(updatedUser), err
}

func (u UserServiceImpl) Delete(id uint) (*domain.UserResponse, error) {
	_, err := u.Repository.Delete(id, time.Now())
	if err != nil {
		return nil, err
	}

	return nil, err
}

func NewUserService(movieRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		Repository: movieRepo,
	}
}

func (u UserServiceImpl) convertToResponse(user *model.User) *domain.UserResponse {
	return &domain.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Address:   user.Address,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u UserServiceImpl) convertToResponses(users *[]model.User) *[]domain.UserResponse {
	var userResponses []domain.UserResponse
	for _, user := range *users {
		userResponses = append(userResponses, domain.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Phone:     user.Phone,
			Address:   user.Address,
			Email:     user.Email,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return &userResponses
}
