package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lcv-back/goload/internal/dataaccess/database"
)

type CreateAccountParams struct {
	AccountName string
	Password    string
}

type CreateAccountOutput struct {
	ID          uint64
	AccountName string
}

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error)
}

type account struct {
	accountDataAccessor         database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor
	hashLogic                   Hash
}

func NewAccount(
	accountDataAccessor database.AccountDataAccessor,
	accountPasswordDataAccessor database.AccountPasswordDataAccessor,
	hashLogic Hash,
) *account {
	return &account{
		accountDataAccessor:         accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic:                   hashLogic,
	}
}

func (a account) isAccountAccountnameTaken(ctx context.Context, accountname string) (bool, error) {
	if _, err := a.accountDataAccessor.GetAccountByAccountname(ctx, accountname); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	accountnameTaken, err := a.isAccountAccountnameTaken(ctx, params.AccountName)

	if err != nil {
		return CreateAccountOutput{}, err
	}

	if accountnameTaken {
		return CreateAccountOutput{}, errors.New("accountname already taken")
	}

	accountID, err := a.accountDataAccessor.CreateAccount(ctx, database.Account{
		Accountname: params.AccountName,
	})

	if err != nil {
		return CreateAccountOutput{}, err
	}

	hashedPassword, err := a.hashLogic.Hash(ctx, params.Password)

	if err != nil {
		return CreateAccountOutput{}, err
	}

	if err := a.accountPasswordDataAccessor.CreateAccountPassword(ctx, database.AccountPassword{
		OfAccountID: accountID,
		Hash:        hashedPassword,
	}); err != nil {
		return CreateAccountOutput{}, err
	}

	return CreateAccountOutput{
		ID:          accountID,
		AccountName: params.AccountName,
	}, nil
}
