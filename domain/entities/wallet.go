package entities

import (
	"e_wallet/utils"
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	Balance   uint      `gorm:"not null"`
	Number    string    `gorm:"not null;unique"`
	User      User      `gorm:"foreignKey:UserID;refer:users"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

func NewWallet(UserID uint) *Wallet {
	return &Wallet{
		UserID:  UserID,
		Balance: 0,
		Number:  utils.GenerateWalletNumber(),
	}
}
