package request

type TopupRequest struct {
	WalletNumber string   `json:"walletNumber"`
	Bank         Bank     `json:"bank,omitempty"`
	Merchant     Merchant `json:"merchant,omitempty"`
	Amount       uint     `json:"amount"`
}

type Bank struct {
	BankName string `json:"bankName"`
	BankCode string `json:"bankCode"`
}

type Merchant struct {
	MerchantName string `json:"merchantName"`
	MerchantCode string `json:"merchantCode"`
}

type PaymentOption struct {
	ID        uint       `json:"id"`
	Method    string     `json:"method"`
	Banks     []Bank     `json:"banks,omitempty"`
	Merchants []Merchant `json:"merchants,omitempty"`
}
