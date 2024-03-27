package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar; not null"`
	Phone     string    `json:"phone" gorm:"type:varchar; not null"`
	Address   string    `json:"address" gorm:"type:varchar; not null"`
	Email     string    `json:"email" gorm:"type:varchar; not null"`
	Username  string    `json:"username" gorm:"type:varchar; not null"`
	Password  string    `json:"password" gorm:"type:varchar; not null"`
	Deleted   bool      `json:"deleted" gorm:"type:boolean; default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}
