package repositories

import (
	"e_wallet/domain/payload/response"
)

type TransactionRepository interface {
	ShowHistoryTransactionByWalletId(walletID uint) (*response.HistoryResponse, error)
}
