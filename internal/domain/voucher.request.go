package domain

import "time"

type VoucherRequest struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Value       int64     `json:"value"`
	Used        bool      `json:"used"`
	ExpiredTime time.Time `json:"expired_time"`
}

type VoucherResponse struct {
	ID          uint         `json:"id"`
	User        UserResponse `json:"user"`
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Value       int64        `json:"value"`
	Used        bool         `json:"used"`
	ExpiredTime time.Time    `json:"expired_time"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
