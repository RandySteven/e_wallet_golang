package apps

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/payload/request"
	"e_wallet/infrastructure/persistences"
	"sync"
)

type userApp struct {
	repo persistences.Repositories
}

// AddUser implements UserAppInterface.
func (app *userApp) AddUser(request *request.UserRequest) (*entities.UserDetail, error) {
	user := entities.UserFromUserRequest(request)
	user, err := app.repo.User.Save(user)
	if err != nil {
		return nil, err
	}
	var (
		profile *entities.UserProfile
		wallet  *entities.Wallet
	)

	var profileErr, walletErr error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		profile = entities.NewUserProfile(user.ID, request)
		profile, profileErr = app.repo.UserProfile.Save(profile)
	}()

	go func() {
		defer wg.Done()
		wallet = entities.NewWallet(user.ID)
		wallet, walletErr = app.repo.Wallet.Save(wallet)
	}()

	wg.Wait()

	if profileErr != nil {
		return nil, profileErr
	}

	if walletErr != nil {
		return nil, walletErr
	}

	userDetail := entities.NewUserDetail(user, profile, wallet)
	return userDetail, nil
}

// GetAllUsers implements UserAppInterface.
func (app *userApp) GetAllUsers() ([]entities.User, error) {
	return app.repo.User.FindAll()
}

// GetUserById implements UserAppInterface.
func (app *userApp) GetUserById(id string) (*entities.UserDetail, error) {
	return app.repo.User.GetUserById(id)
}

func NewUserApp(repo persistences.Repositories) *userApp {
	return &userApp{repo}
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	AddUser(request *request.UserRequest) (*entities.UserDetail, error)
	GetAllUsers() ([]entities.User, error)
	GetUserById(id string) (*entities.UserDetail, error)
}
