// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wiring

import (
	"github.com/google/wire"
	"github.com/lcv-back/goload/internal/configs"
	"github.com/lcv-back/goload/internal/dataaccess"
	"github.com/lcv-back/goload/internal/dataaccess/database"
	"github.com/lcv-back/goload/internal/handler"
	"github.com/lcv-back/goload/internal/handler/grpc"
	"github.com/lcv-back/goload/internal/logic"
)

// Injectors from wire.go:

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	config, err := configs.NewConfig(configFilePath)
	if err != nil {
		return nil, nil, err
	}
	configsDatabase := config.Database
	db, cleanup, err := database.InitializeDB(configsDatabase)
	if err != nil {
		return nil, nil, err
	}
	goquDatabase := database.InitializeGoquDB(db)
	accountDataAccessor := database.NewAccountDataAccessor(goquDatabase)
	accountPasswordDataAccessor := database.NewAccountPasswordDataAccessor(goquDatabase)
	account := config.Account
	hash := logic.NewHash(account)
	logicAccount := logic.NewAccount(goquDatabase, accountDataAccessor, accountPasswordDataAccessor, hash)
	goLoadServiceServer := grpc.NewHandler(logicAccount)
	server := grpc.NewServer(goLoadServiceServer)
	return server, func() {
		cleanup()
	}, nil
}

// wire.go:

var WireSet = wire.NewSet(configs.WireSet, dataaccess.WireSet, logic.WireSet, handler.WireSet)
