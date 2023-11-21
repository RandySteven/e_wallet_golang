package persistences

import (
	"e_wallet/domain/payload/response"
	"e_wallet/domain/repositories"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

// ShowHistoryTransactionByWalletId implements repositories.TransactionRepository.
func (repo *TransactionRepository) ShowHistoryTransactionByWalletId(walletID uint) (*response.HistoryResponse, error) {
	var histories []response.HistoryResponse
	err := repo.db.Raw("? UNION ? UNION ?",
		repo.db.Raw("SELECT u2.name AS user_name, amount, transfer_transactions.created_at, CONCAT('Receive money from ', users.name) AS description FROM transfer_transactions  JOIN wallets ON transfer_transactions.sender_wallet_id = wallets.id JOIN users ON  wallets.user_id = users.id JOIN wallets w2 ON transfer_transactions.receiver_wallet_id = w2.id JOIN users u2 ON w2.user_id = u2.id WHERE receiver_wallet_id = 1 AND transfer_transactions.created_at BETWEEN '2022-12-31 23:59:59' AND '2024-01-01 00:00:00' GROUP BY u2.name, amount, transfer_transactions.created_at, users.name"),
		repo.db.Raw("SELECT u2.name AS user_name, amount, transfer_transactions.created_at, CONCAT('Send into ', users.name) AS description FROM transfer_transactions JOIN wallets ON transfer_transactions.receiver_wallet_id = wallets.id JOIN users ON wallets.user_id = users.id JOIN wallets w2 ON transfer_transactions.sender_wallet_id = w2.id JOIN users u2 ON w2.user_id = u2.id WHERE sender_wallet_id = 1 AND transfer_transactions.created_at BETWEEN '2022-12-31 23:59:59' AND '2024-01-01 00:00:00' GROUP BY u2.name, amount, transfer_transactions.created_at, users.name"),
		repo.db.Raw("SELECT users.name AS user_name, amount, topup_transactions.created_at, CONCAT('Top up from ', source_of_fund) AS description FROM topup_transactions JOIN wallets ON topup_transactions.wallet_id = wallets.id JOIN users ON wallets.user_id = users.id WHERE wallet_id = 1 AND topup_transactions.created_at BETWEEN '2022-12-31 23:59:59' AND '2024-01-01 00:00:00' GROUP BY users.name, amount, topup_transactions.created_at, source_of_fund"),
	).Order("created_at").Scan(&histories).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

var _ repositories.TransactionRepository = &TransactionRepository{}
