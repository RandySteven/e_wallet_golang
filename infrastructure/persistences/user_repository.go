package persistences

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// FindAll implements repositories.UserRepository.
func (repo *UserRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := repo.db.Table("users").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById implements repositories.UserRepository.
func (repo *UserRepository) GetUserById(id string) (*entities.UserDetail, error) {
	var userDetail *entities.UserDetail
	err := repo.db.Table("wallets").
		Preload("Users").Scan(&userDetail).Error
	if err != nil {
		return nil, err
	}
	return userDetail, nil
}

// Save implements repositories.UserRepository.
func (repo *UserRepository) Save(user *entities.User) (*entities.User, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

var _ repositories.UserRepository = &UserRepository{}
