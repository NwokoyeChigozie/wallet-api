package model

import (
	"time"
)

type User struct {
	ID           uint      `json:"id"`
	FirstName    string    `gorm:"column:first_name; type:text; not null" json:"first_name"`
	LastName     string    `gorm:"column:last_name; type:text;not null" json:"last_name"`
	Email        string    `gorm:"column:email; type:text;size:191;not null" json:"email"`
	Password     string    `gorm:"column:password; type:text; not null" json:"-"`
	Token        string    `gorm:"column:token; type:text; null" json:"token"`
	TokenExpires time.Time `gorm:"column:token_expires; null; autoCreateTime" json:"token_expires"`
	WalletID     uint      `gorm:"-:all" json:"wallet_id"`
	CreatedAt    time.Time `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
}

type CreateUserRequestModel struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type LoginRequestModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserIdentity struct {
	ID int `json:"id"`
}
