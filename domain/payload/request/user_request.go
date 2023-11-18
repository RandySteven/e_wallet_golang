package request

import "time"

type UserRequest struct {
	Name     string    `json:"name"`
	UserName string    `json:"userName"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
