package entities

import (
	"e_wallet/domain/payload/request"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Birthday  time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

type UserDetail struct {
	Name     string
	Email    string
	Birthday time.Time
	Wallet   Wallet
}

func UserFromUserRequest(request *request.UserRequest) *User {
	return &User{
		Name:     request.Name,
		Birthday: request.Birthday,
	}
}

func NewUserDetail(user *User, profile *UserProfile, wallet *Wallet) *UserDetail {
	return &UserDetail{
		Name:     user.Name,
		Email:    profile.Email,
		Birthday: user.Birthday,
		Wallet:   *wallet,
	}
}
