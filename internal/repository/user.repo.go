package repository

import (
	"errors"
	"gorm.io/gorm"
	"royalty-api/internal/domain"
	"royalty-api/internal/model"
	"strconv"
	"time"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) List() (*[]model.User, error) {
	var users []model.User

	db := u.db.Model(&users)

	checkUsers := db.Where("deleted = ? OR deleted is null", false).Find(&users)
	if checkUsers.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &users, nil
}

func (u userRepository) View(id uint) (*model.User, error) {
	var user model.User

	db := u.db.Model(&user)

	checkUsers := db.Debug().Where("deleted = ? OR deleted is null", false).First(&user)
	if checkUsers.Error != nil && checkUsers.Error.Error() != "record not found" {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	if user.ID < 1 {
		return nil, errors.New("User not found ")
	}

	return &user, nil
}

func (u userRepository) Insert(input *domain.UserRequest, timeNow time.Time) (*model.User, error) {
	var user model.User

	db := u.db.Model(&user)

	user = model.User{
		Name:     input.Name,
		Phone:    input.Phone,
		Address:  input.Address,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
		Deleted:  false,
	}

	addUser := db.Debug().Create(&user)
	if addUser.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &user, nil
}

func (u userRepository) Update(input *domain.UserRequest, timeNow time.Time) (*model.User, error) {
	var user model.User

	db := u.db.Model(&user)

	user.ID = input.ID

	checkUser := db.Debug().First(&user, "deleted=false")
	if checkUser.RowsAffected < 1 {
		return nil, errors.New("User with ID " + strconv.Itoa(int(input.ID)) + " not found ")
	}

	user.Name = input.Name
	user.Phone = input.Phone
	user.Address = input.Address
	user.Email = input.Email
	user.Username = input.Username
	user.Password = input.Password
	user.UpdatedAt = timeNow

	addUser := db.Debug().Updates(&user)
	if addUser.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return &user, nil
}

func (u userRepository) Delete(id uint, timeNow time.Time) (*model.User, error) {
	var user model.User

	db := u.db.Model(&user)

	user.ID = id

	checkUser := db.Debug().First(&user, "deleted=false")
	if checkUser.RowsAffected < 1 {
		return nil, errors.New("User with ID " + strconv.Itoa(int(id)) + " not found ")
	}

	user.Deleted = true
	user.UpdatedAt = timeNow

	deleteUser := db.Debug().Updates(&user)
	if deleteUser.Error != nil {
		return nil, errors.New("Internal Server Error, please contact our customer service ")
	}

	return nil, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
