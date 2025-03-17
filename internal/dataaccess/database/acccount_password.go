package database

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type AccountPassword struct {
	OfAccountID uint64 `sql:"of_account_id"`
	Hash        string `sql:"hash"`
}

type AccountPasswordDataAccessor interface {
	CreateAccountPassword(ctx context.Context, accountPass AccountPassword) error
	GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error)
	UpdateAccountPassword(ctx context.Context, accountPass AccountPassword) error
	WithDatabase(database Database) AccountPasswordDataAccessor
}

type accountPasswordDataAccessor struct {
	database Database
}

func NewAccountPasswordDataAccessor(database *goqu.Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
	}
}

func (a *accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPass AccountPassword) error {
	panic("unimplemented")
}

func (a *accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPass AccountPassword) error {
	panic("unimplemented")
}

func (a *accountPasswordDataAccessor) GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error) {
	var accountPassword AccountPassword
	found, err := a.database.From("account_passwords").
		Where(goqu.Ex{"of_account_id": ofAccountID}).
		ScanStructContext(ctx, &accountPassword)

	if err != nil {
		return AccountPassword{}, err
	}

	if !found {
		return AccountPassword{}, sql.ErrNoRows
	}

	return accountPassword, nil
}

func (a accountPasswordDataAccessor) WithDatabase(database Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
	}
}
