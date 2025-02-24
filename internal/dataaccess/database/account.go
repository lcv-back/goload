package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type Account struct {
	UserID   uint64 `sql:"user_id"`
	Username string `sql:"username"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	GetAccountByID(ctx context.Context, id uint64) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
}

type accountDataAccessor struct {
	database *goqu.Database
}

func NewAccountDataAccessor(database *goqu.Database) AccountDataAccessor {
	return &accountDataAccessor{database: database}
}

func (a accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64, error) {
	result, err := a.database.
		Insert("accounts").
		Rows(goqu.Record{
			"username": account.Username,
		}).
		Executor().
		ExecContext(ctx)

	if err != nil {
		log.Printf("failed to create account, err= %+v", err)
	}

	LastInsertedId, err := result.LastInsertId()
	if err != nil {
		log.Printf("failed to get last inserted id, err= %+v", err)
		return 0, err
	}

	return uint64(LastInsertedId), nil
}

func (a *accountDataAccessor) GetAccountByID(ctx context.Context, id uint64) (Account, error) {
	// Implement the method
	return Account{}, nil
}

func (a *accountDataAccessor) GetAccountByUsername(ctx context.Context, username string) (Account, error) {
	// Implement the method
	return Account{}, nil
}
