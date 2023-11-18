package repositories

import "e_wallet/domain/entities"

type UserProfileRepository interface {
	Save(profile *entities.UserProfile) (*entities.UserProfile, error)
}
