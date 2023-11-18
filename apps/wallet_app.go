package apps

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"
)

type walletApp struct {
	repo repositories.WalletRepository
}

// CreateWallet implements WalletAppInteface.
func (app *walletApp) CreateWallet(userId uint) (*entities.Wallet, error) {
	wallet := entities.NewWallet(userId)
	return app.repo.Save(wallet)
}

func NewWalletApp(repo repositories.WalletRepository) *walletApp {
	return &walletApp{repo}
}

var _ WalletAppInteface = &walletApp{}

type WalletAppInteface interface {
	CreateWallet(userId uint) (*entities.Wallet, error)
}
