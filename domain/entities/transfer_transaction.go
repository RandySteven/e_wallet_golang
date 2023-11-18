package entities

import (
	"time"

	"gorm.io/gorm"
)

type TransferTransaction struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	SenderWalletID   uint      `gorm:"not null"`
	ReceiverWalletID uint      `gorm:"not null"`
	Amount           uint      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt        gorm.DeletedAt
	SenderWallet     Wallet `gorm:"foreignKey:SenderWalletID;refer:wallets"`
	ReceiverWallet   Wallet `gorm:"foreignKey:ReceiverWalletID;refer:wallets"`
}
