package response

import "e_wallet/domain/entities"

type TransferResponse struct {
	Sender   entities.Wallet
	Receiver entities.Wallet
}
