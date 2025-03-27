package mysql

import (
	"context"
	"database/sql"

	"go.uber.org/zap"

	"github.com/doug-martin/goqu/v9"

	db "github.com/lcv-back/goload/internal/dataaccess/database"
	"github.com/lcv-back/goload/internal/utils"
)

const (
	TabNameTokenPublicKeys          = "token_public_keys"
	ColNameTokenPublicKeysID        = "id"
	ColNameTokenPublicKeysPublicKey = "public_key"
)

type TokenPublicKey struct {
	ID        uint64 `sql:"id"`
	PublicKey []byte `sql:"public_key"`
}

type TokenPublicKeyDataAccessor interface {
	CreatePublicKey(ctx context.Context, tokenPublicKey TokenPublicKey) (uint64, error)
	GetPublicKey(ctx context.Context, id uint64) (TokenPublicKey, error)
	WithDatabase(database db.Database) TokenPublicKeyDataAccessor
}

type tokenPublicKeyDataAccessor struct {
	database db.Database
	logger   *zap.Logger
}

func NewTokenPublicKeyDataAccessor(
	database *goqu.Database,
	logger *zap.Logger,
) TokenPublicKeyDataAccessor {
	return &tokenPublicKeyDataAccessor{
		database: database,
		logger:   logger,
	}
}

func (a tokenPublicKeyDataAccessor) CreatePublicKey(ctx context.Context, tokenPublicKey TokenPublicKey) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, a.logger)
	result, err := a.database.
		Insert(TabNameTokenPublicKeys).
		Rows(goqu.Record{
			ColNameTokenPublicKeysPublicKey: tokenPublicKey.PublicKey,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create token public key")
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get last inserted id")
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (a tokenPublicKeyDataAccessor) GetPublicKey(ctx context.Context, id uint64) (TokenPublicKey, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Uint64("id", id))

	tokenPublicKey := TokenPublicKey{}
	found, err := a.database.Select().From(TabNameTokenPublicKeys).Where(goqu.Ex{
		ColNameTokenPublicKeysID: id,
	}).
		Executor().
		ScanStructContext(ctx, &tokenPublicKey)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get public key")
		return TokenPublicKey{}, err
	}

	if !found {
		logger.Warn("public key not found")
		return TokenPublicKey{}, sql.ErrNoRows
	}

	return tokenPublicKey, nil
}

func (a tokenPublicKeyDataAccessor) WithDatabase(database db.Database) TokenPublicKeyDataAccessor {
	a.database = database
	return a
}
