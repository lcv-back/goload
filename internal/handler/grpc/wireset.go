package grpc

import (
	"github.com/google/wire"
	"github.com/lcv-back/goload/internal/logic"
)

var WireSet = wire.NewSet(
	NewHandler,
	NewServer,
	logic.WireSet,
)
