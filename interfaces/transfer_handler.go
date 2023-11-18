package interfaces

import (
	"e_wallet/apps"
	"e_wallet/domain/payload/request"
	"e_wallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferTransaction struct {
	transferApp apps.TransferAppInterface
}

func NewTransferTransaction(transferApp apps.TransferAppInterface) *TransferTransaction {
	return &TransferTransaction{transferApp: transferApp}
}

func (transferTransaction *TransferTransaction) TransferTransaction(c *gin.Context) {
	var request *request.TransferRequest
	var rh utils.ResponseHandler
	if err := c.ShouldBind(&request); err != nil {
		rh.ResponseEncoder(c, http.StatusBadRequest, false, "message", "Bad request")
		return
	}
	transfer, err := transferTransaction.transferApp.TransferTransaction(request)
	if err != nil {
		rh.ResponseEncoder(c, http.StatusInternalServerError, false, "message", "Internal server error")
		return
	}
	rh.ResponseEncoder(c, http.StatusCreated, true, "transaction", transfer)
}
