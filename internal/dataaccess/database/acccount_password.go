package database

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/lcv-back/goload/internal/utils"
	"go.uber.org/zap"
)

const (
	TabNameAccountPasswords            = "account_passwords"
	ColNameAccountPasswordsOfAccountID = "of_account_id"
	ColNameAccountPasswordsHash        = "hash"
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
	logger   *zap.Logger
}

func NewAccountPasswordDataAccessor(
	database *goqu.Database,
	logger *zap.Logger,
) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
		logger:   logger,
	}
}

func (a accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	logger := utils.LoggerWithContext(ctx, a.logger)
	_, err := a.database.
		Insert(TabNameAccountPasswords).
		Rows(goqu.Record{
			ColNameAccountPasswordsOfAccountID: accountPassword.OfAccountID,
			ColNameAccountPasswordsHash:        accountPassword.Hash,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create account password")
		return err
	}

	return nil
}

func (a accountPasswordDataAccessor) GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Uint64("of_account_id", ofAccountID))
	accountPassword := AccountPassword{}
	found, err := a.database.
		From(TabNameAccounts).
		Where(goqu.Ex{ColNameAccountPasswordsOfAccountID: ofAccountID}).
		ScanStructContext(ctx, &accountPassword)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account password by id")
		return AccountPassword{}, err
	}

	if !found {
		logger.Warn("cannot find account by id")
		return AccountPassword{}, sql.ErrNoRows
	}

	return accountPassword, nil
}

func (a accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	logger := utils.LoggerWithContext(ctx, a.logger)
	_, err := a.database.
		Update(TabNameAccountPasswords).
		Set(goqu.Record{ColNameAccountPasswordsHash: accountPassword.Hash}).
		Where(goqu.Ex{ColNameAccountPasswordsOfAccountID: accountPassword.OfAccountID}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to update account password")
		return err
	}

	return nil
}

func (a accountPasswordDataAccessor) WithDatabase(database Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
	}
}
