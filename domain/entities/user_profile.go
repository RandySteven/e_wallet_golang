package entities

import (
	"e_wallet/domain/payload/request"
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
	User      User `gorm:"foreignKey:UserID;refer:users"`
}

func NewUserProfile(userId uint, request *request.UserRequest) *UserProfile {
	return &UserProfile{
		Name:     request.UserName,
		Email:    request.Email,
		Password: request.Password,
		UserID:   userId,
	}
}
