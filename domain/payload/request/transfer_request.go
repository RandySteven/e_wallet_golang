package request

type TransferRequest struct {
	SenderWalletNumber   string `json:"senderWalletNumber"`
	ReceiverWalletNumber string `json:"receiverWalletNumber"`
	Amount               uint   `json:"amount"`
}
