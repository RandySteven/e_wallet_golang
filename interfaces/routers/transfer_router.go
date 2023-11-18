package routers

import (
	"e_wallet/interfaces"

	"github.com/gin-gonic/gin"
)

func TransferRouter(r *gin.RouterGroup, transfer *interfaces.TransferTransaction) {
	transferRouter := r.Group("/transfers")
	transferRouter.POST("/create-transfer", transfer.TransferTransaction)
}
