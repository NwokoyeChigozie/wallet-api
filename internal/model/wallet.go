package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID            uint            `json:"id"`
	UserID        uint            `gorm:"column:user_id; type:uint; not null" json:"user_id"`
	Balance       string          `gorm:"column:balance; type:text; not null" json:"-"`
	WalletBalance decimal.Decimal `gorm:"-:all" json:"wallet_balance"`
	CreatedAt     time.Time       `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
}

type Transactions struct {
	ID        uint      `json:"id"`
	UserID    uint      `gorm:"column:user_id; type:uint; not null" json:"user_id"`
	Type      string    `gorm:"column:type; type:text; not null" json:"type"`
	Amount    string    `gorm:"column:amount; type:text; not null" json:"amount"`
	CreatedAt time.Time `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
}

type DebitOrCreditWalletRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}
