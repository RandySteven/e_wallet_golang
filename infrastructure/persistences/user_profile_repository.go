package persistences

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"

	"gorm.io/gorm"
)

type UserProfile struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) *UserProfile {
	return &UserProfile{db}
}

// Save implements repositories.UserProfileRepository.
func (repo *UserProfile) Save(profile *entities.UserProfile) (*entities.UserProfile, error) {
	err := repo.db.Create(&profile).Error
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func NewUserProfile(db *gorm.DB) *UserProfile {
	return &UserProfile{db}
}

var _ repositories.UserProfileRepository = &UserProfile{}
