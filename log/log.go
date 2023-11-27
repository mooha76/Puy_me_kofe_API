package boot

import (
	"fmt"
	"log/slog"
	"time"

	elPath "github.com/Junvary/erleng/path"
	"github.com/mooha76/Puy_me_kofe_API/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitializeLogger(cfg *config.Logger) *slog.Logger {

	if err := elPath.CreatePath(cfg.Path); err != nil {
		panic(err)
	}
	fmt.Println("Successfully Loger initailsed")
	lj := &lumberjack.Logger{
		Filename:   cfg.Path + "/" + cfg.Filename + "-" + time.Now().Format("20060102150405") + ".log",
		LocalTime:  true,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
	}
	defer lj.Close()
	sl := slog.NewJSONHandler(lj, nil)
	logger := slog.New(sl)
	return logger
}
