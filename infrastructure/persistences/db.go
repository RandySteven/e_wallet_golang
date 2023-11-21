package persistences

import (
	"e_wallet/domain/entities"
	"e_wallet/domain/repositories"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User                repositories.UserRepository
	UserProfile         repositories.UserProfileRepository
	Wallet              repositories.WalletRepository
	TransferTransaction repositories.TransferRepository
	TopupTransaction    repositories.TopupRepository
	db                  *gorm.DB
}

func NewRepository(dbHost, dbName, dbUser, dbPass, dbPort string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return &Repositories{
		User:                NewUserRepository(db),
		UserProfile:         NewUserProfileRepository(db),
		Wallet:              NewWalletRepository(db),
		TransferTransaction: NewTransferRepository(db),
		TopupTransaction:    NewTopupRepository(db),
		db:                  db,
	}, nil
}

func (r *Repositories) Close() <-chan struct{} {
	return r.db.Statement.Context.Done()
}

func (r *Repositories) Automigrate() error {
	return r.db.AutoMigrate(
		&entities.User{},
		&entities.UserProfile{},
		&entities.Wallet{},
		&entities.TransferTransaction{},
		&entities.TopupTransaction{},
	)
}
