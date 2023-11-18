package repositories

import "e_wallet/domain/entities"

type WalletRepository interface {
	Save(wallet *entities.Wallet) (*entities.Wallet, error)
	GetWalletByWalletNumber(walletNumber string) (*entities.Wallet, error)
}
