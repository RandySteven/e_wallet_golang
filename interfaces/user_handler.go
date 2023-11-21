package interfaces

import (
	"e_wallet/apps"
	"e_wallet/domain/payload/request"
	"e_wallet/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	validate *validator.Validate
	userApp  apps.UserAppInterface
}

func NewUser(userApp apps.UserAppInterface, validate *validator.Validate) *User {
	return &User{userApp: userApp, validate: validate}
}

func (u *User) RegisterUser(c *gin.Context) {
	var request *request.UserRequest
	var rh utils.ResponseHandler
	if err := c.ShouldBind(&request); err != nil {
		rh.ResponseEncoder(c, http.StatusBadRequest, false, "message", "Bad request")
		return
	}
	err := u.validate.Struct(request)
	errs := strings.Split(err.Error(), "\n")
	if errs != nil {
		rh.ResponseEncoder(c, http.StatusBadRequest, false, "errors", errs)
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
