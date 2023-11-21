package interfaces

import (
	"e_wallet/apps"

	"github.com/gin-gonic/gin"
)

type Topup struct {
	app apps.TopupAppInterface
}

func NewTopup(app apps.TopupAppInterface) *Topup {
	return &Topup{app}
}

func (topup *Topup) TopupTransaction(c *gin.Context) {

}
