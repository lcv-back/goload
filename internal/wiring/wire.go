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
)

var WireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
