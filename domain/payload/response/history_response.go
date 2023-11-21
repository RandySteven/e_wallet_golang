package response

import "time"

type HistoryResponse struct {
	UserName    string
	Amount      uint
	CreatedAt   time.Time
	Description string
}
