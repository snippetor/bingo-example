package log

import (
	"time"
	"github.com/snippetor/bingo/log"
)

var Logger = log.NewLoggerWithConfig(&log.Config{
	Level:                  log.Info,
	OutputType:             log.Console | log.File,
	LogFileRollingType:     log.RollingDaily,
	LogFileOutputDir:       ".",
	LogFileName:            "echo",
	LogFileNameDatePattern: "20060102",
	LogFileNameExt:         ".log",
	LogFileMaxSize:         500 * log.MB,
	LogFileScanInterval:    1 * time.Second,
})
