package persistences

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

// GetWalletByWalletNumber implements repositories.WalletRepository.
func (repo *WalletRepository) GetWalletByWalletNumber(walletNumber string) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := repo.db.
		Table("wallets").
		Where("number = ?", walletNumber).
		Scan(&wallet).
		Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Save implements repositories.WalletRepository.
func (repo *WalletRepository) Save(wallet *entities.Wallet) (*entities.Wallet, error) {
	err := repo.db.Create(&wallet).Error
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db}
}

var _ repositories.WalletRepository = &WalletRepository{}
