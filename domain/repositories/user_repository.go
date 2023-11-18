package repositories

import "e_wallet/domain/entities"

type UserRepository interface {
	Save(user *entities.User) (*entities.User, error)
	FindAll() ([]entities.User, error)
	GetUserById(id string) (*entities.UserDetail, error)
}
