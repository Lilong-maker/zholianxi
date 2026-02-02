package init

import (
	"context"

	l "github.com/gospacex/gospacex/core/logger"
	"github.com/gospacex/gospacex/core/storage/conf"
)

var Log = conf.LogConfig{
	Level:         "info",
	Path:          ".",
	MaxSize:       40,
	MaxBackups:    60,
	MaxAge:        20,
	Compress:      true,
	Format:        "",
	ConsoleEnable: true,
}

func LogInit() {
	l.Init(context.Background(), &Log)
}
