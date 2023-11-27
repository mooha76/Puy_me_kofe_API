package global

import (
	"log/slog"

	"github.com/mooha76/Puy_me_kofe_API/config"
)

var (
	MainConfig config.Config
	MainLogger *slog.Logger
)
