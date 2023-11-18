package entities

import (
	"time"

	"gorm.io/gorm"
)

type TopupTransaction struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	WalletID     uint      `gorm:"not null"`
	SourceOfFund string    `gorm:"not null"`
	Amount       uint      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt
	Wallet       Wallet `gorm:"foreignKey:WalletID;refer:wallets"`
}
