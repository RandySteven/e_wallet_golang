package routers

import (
	"e_wallet/interfaces"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, users *interfaces.User) {
	userGroup := r.Group("/users")
	userGroup.POST("/register", users.RegisterUser)
}
