package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type AccountPassword struct {
	OfUserID uint64 `sql:"of_user_id"`
	Hash     string `sql:"hash"`
}

type AccountPasswordDataAccessor interface {
	CreateUserPassword(ctx context.Context, accountPass AccountPassword) error
	UpdateUserPassword(ctx context.Context, accountPass AccountPassword) error
}

type accountPasswordDataAccessor struct {
	database *goqu.Database
}

func NewAccountPasswordDataAccessor(database *goqu.Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
	}
}

func (a *accountPasswordDataAccessor) CreateUserPassword(ctx context.Context, accountPass AccountPassword) error {
	panic("unimplemented")
}

func (a *accountPasswordDataAccessor) UpdateUserPassword(ctx context.Context, accountPass AccountPassword) error {
	panic("unimplemented")
}
