package persistences

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"

	"gorm.io/gorm"
)

type TopupRepository struct {
	db *gorm.DB
}

// TopupTransaction implements repositories.TopupRepository.
func (repo *TopupRepository) TopupTransaction(topup *entities.TopupTransaction) (*entities.TopupTransaction, error) {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var wallet *entities.Wallet
	err := tx.Model(&entities.Wallet{}).
		Where("id = ?", topup.WalletID).
		Scan(&wallet).
		Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Table("wallets").
		Where("id = ?", wallet.ID).
		Update("balance", gorm.Expr("balance + ?", topup.Amount)).
		Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Model(&entities.TopupTransaction{}).
		Create(&topup).
		Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return topup, nil
}

func NewTopupRepository(db *gorm.DB) *TopupRepository {
	return &TopupRepository{db}
}

var _ repositories.TopupRepository = &TopupRepository{}
