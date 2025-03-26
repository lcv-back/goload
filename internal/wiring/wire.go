//go:build wireinject
// +build wireinject

//
// go:generate go run github.com/google/wire/cmd/wire

package wiring

import (
	"github.com/google/wire"
	"github.com/lcv-back/goload/internal/configs"
	"github.com/lcv-back/goload/internal/dataaccess"
	"github.com/lcv-back/goload/internal/handler"
	"github.com/lcv-back/goload/internal/handler/grpc"
	"github.com/lcv-back/goload/internal/logic"
	"go.uber.org/zap"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
	NewLogger,
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
