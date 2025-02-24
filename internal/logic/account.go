package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lcv-back/goload/internal/dataaccess/database"
)

type CreateAccountParams struct {
	Username  string
	Passoword string
}

type User struct {
	ID       uint64
	Username string
}

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (User, error)
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
) Account {
	return &account{
		accountDataAccessor:         accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic:                   hashLogic,
	}
}

func (a account) isAccountUsernameTaken(ctx context.Context, username string) (bool, error) {
	if _, err := a.accountDataAccessor.GetAccountByUsername(ctx, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (User, error) {
	usernameTaken, err := a.isAccountUsernameTaken(ctx, params.Username)

	if err != nil {
		return User{}, err
	}

	if usernameTaken {
		return User{}, errors.New("username already taken")
	}

	accountID, err := a.accountDataAccessor.CreateAccount(ctx, database.Account{
		Username: params.Username,
	})

	if err != nil {
		return User{}, err
	}

	hashedPassword, err := a.hashLogic.Hash(ctx, params.Passoword)

	if err != nil {
		return User{}, err
	}

	if err := a.accountPasswordDataAccessor.CreateUserPassword(ctx, database.AccountPassword{
		OfUserID: accountID,
		Hash:     hashedPassword,
	}); err != nil {
		return User{}, err
	}

	return User{
		ID:       accountID,
		Username: params.Username,
	}, nil
}
