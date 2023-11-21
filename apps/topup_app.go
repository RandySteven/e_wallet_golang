package apps

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/payload/request"
	"e_wallet/infrastructure/persistences"
)

type topupApp struct {
	repo persistences.Repositories
}

// TopupTransaction implements TopupAppInterface.
func (app *topupApp) TopupTransaction(request *request.TopupRequest) (*entities.TopupTransaction, error) {
	wallet, err := app.repo.Wallet.GetWalletByWalletNumber(request.WalletNumber)
	if err != nil {
		return nil, err
	}
	topup := &entities.TopupTransaction{
		WalletID: wallet.ID,
		Amount:   request.Amount,
	}
	return app.repo.TopupTransaction.TopupTransaction(topup)
}

func NewTopup(repo persistences.Repositories) *topupApp {
	return &topupApp{repo: repo}
}

var _ TopupAppInterface = &topupApp{}

type TopupAppInterface interface {
	TopupTransaction(request *request.TopupRequest) (*entities.TopupTransaction, error)
}
