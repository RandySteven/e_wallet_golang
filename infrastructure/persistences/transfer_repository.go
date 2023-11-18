package persistences

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"

	"gorm.io/gorm"
)

type TransferRepository struct {
	db *gorm.DB
}

// TransferTransaction implements repositories.TransferRepository.
func (repo *TransferRepository) TransferTransaction(transfer *entities.TransferTransaction) (*entities.TransferTransaction, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var senderWallet *entities.Wallet
	err := tx.Model(&entities.Wallet{}).
		Where("id = ?", transfer.SenderWalletID).
		Scan(&senderWallet).Error
	if err != nil {
		return nil, err
	}

	if senderWallet.Balance < transfer.Amount {
		tx.Rollback()
		return nil, err
	}

	err = tx.Table("wallets").
		Where("id = ?", senderWallet.ID).
		Update("balance", gorm.Expr("balance - ?", transfer.Amount)).
		Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Table("wallets").
		Where("id = ?", transfer.ReceiverWalletID).
		Update("balance", gorm.Expr("balance + ?", transfer.Amount)).
		Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Create(&transfer).Error
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return transfer, nil
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	return &TransferRepository{db}
}

var _ repositories.TransferRepository = &TransferRepository{}
