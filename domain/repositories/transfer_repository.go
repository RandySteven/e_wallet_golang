package repositories

import "e_wallet/domain/entities"

type TransferRepository interface {
	TransferTransaction(transfer *entities.TransferTransaction) (*entities.TransferTransaction, error)
}
