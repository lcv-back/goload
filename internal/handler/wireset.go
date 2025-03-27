package handler

import (
	"github.com/google/wire"
	"github.com/lcv-back/goload/internal/handler/grpc"
	"github.com/lcv-back/goload/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)
