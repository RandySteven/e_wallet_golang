package request

import "time"

type UserRequest struct {
	Name     string    `json:"name" validate:"required,min=3,max=32"`
	UserName string    `json:"userName"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password"`
}
