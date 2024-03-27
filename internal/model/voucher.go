package model

import (
	"gorm.io/gorm"
	"time"
)

type Voucher struct {
	ID          uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserID      uint      `json:"-"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string    `json:"name" gorm:"type:varchar; not null"`
	Code        string    `json:"code" gorm:"type:varchar; not null"`
	Value       int64     `json:"value" gorm:"type:bigint; not null"`
	Used        bool      `json:"used"`
	ExpiredTime time.Time `json:"expired_time" gorm:"not null"`
	Deleted     bool      `json:"deleted" gorm:"type:boolean; default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *Voucher) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}
