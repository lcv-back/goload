package dataaccess

import (
	"github.com/google/wire"
	"github.com/lcv-back/goload/internal/dataaccess/database"
)

var WireSet = wire.NewSet(
	database.WireSet,
)
