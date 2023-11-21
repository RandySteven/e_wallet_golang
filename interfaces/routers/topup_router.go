package routers

import (
	"e_wallet/interfaces"

	"github.com/gin-gonic/gin"
)

func TopupRouter(r *gin.RouterGroup, topup *interfaces.Topup) {
	topupRouter := r.Group("/topup")
	topupRouter.POST("", topup.TopupTransaction)
}
