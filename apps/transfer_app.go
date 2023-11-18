package apps

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/payload/request"
	"e_wallet/infrastructure/persistences"
)

type transferApp struct {
	repo persistences.Repositories
}

// TransferTransaction implements TransferAppInterface.
func (app *transferApp) TransferTransaction(request *request.TransferRequest) (*entities.TransferTransaction, error) {
	var transfer *entities.TransferTransaction
	sender, err := app.repo.Wallet.GetWalletByWalletNumber(request.SenderWalletNumber)
	if err != nil {
		return nil, err
	}
	receiver, err := app.repo.Wallet.GetWalletByWalletNumber(request.ReceiverWalletNumber)
	if err != nil {
		return nil, err
	}

	transfer = &entities.TransferTransaction{
		SenderWalletID:   sender.ID,
		ReceiverWalletID: receiver.ID,
		Amount:           request.Amount,
	}
	return app.repo.TransferTransaction.TransferTransaction(transfer)
}

func NewTransferApp(repo persistences.Repositories) *transferApp {
	return &transferApp{repo}
}

var _ TransferAppInterface = &transferApp{}

type TransferAppInterface interface {
	TransferTransaction(request *request.TransferRequest) (*entities.TransferTransaction, error)
}
