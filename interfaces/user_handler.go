package interfaces

import (
	"e_wallet/apps"
	"e_wallet/domain/payload/request"
	"e_wallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userApp apps.UserAppInterface
}

func NewUser(userApp apps.UserAppInterface) *User {
	return &User{userApp: userApp}
}

func (u *User) RegisterUser(c *gin.Context) {
	var request *request.UserRequest
	var rh utils.ResponseHandler
	if err := c.ShouldBind(&request); err != nil {
		rh.ResponseEncoder(c, http.StatusBadRequest, false, "message", "Bad request")
		return
	}
	userDetail, err := u.userApp.AddUser(request)
	if err != nil {
		rh.ResponseEncoder(c, http.StatusInternalServerError, false, "message", "Failed to create user")
		return
	}
	rh.ResponseEncoder(c, http.StatusCreated, true, "user", userDetail)
}

func (u *User) GetAllUsers(c *gin.Context) {

}

func (u *User) GetUserById(c *gin.Context) {}
