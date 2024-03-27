package model

import (
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID        uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar"`
	Deleted   bool      `json:"deleted" gorm:"type:boolean; default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}
