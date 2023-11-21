package repositories

import "e_wallet/domain/entities"

type TopupRepository interface {
	TopupTransaction(topup *entities.TopupTransaction) (*entities.TopupTransaction, error)
}
