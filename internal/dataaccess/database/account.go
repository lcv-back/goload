package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type Account struct {
	AccountID   uint64 `sql:"account_id"`
	Accountname string `sql:"accountname"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	GetAccountByID(ctx context.Context, id uint64) (Account, error)
	GetAccountByAccountname(ctx context.Context, accountname string) (Account, error)
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
			"accountname": account.Accountname,
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

func (a *accountDataAccessor) GetAccountByAccountname(ctx context.Context, accountname string) (Account, error) {
	// Implement the method
	return Account{}, nil
}
