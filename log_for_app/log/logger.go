package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"log"
)

type level int

var infoLevel level = 1
var debugLevel level = 2
var errorLevel level = 3

type Logger struct {
	level   level
	zeroLog zerolog.Logger
}

var logger *Logger

func init() {
	zeroLogger := zerolog.New(log.Writer()).With().Timestamp().Str("tag", "go-log").Logger()
	logger = &Logger{
		infoLevel,
		zeroLogger,
	}
}

func Print(msg ...interface{}) {
	if logger.level <= 1 {
		logger.zeroLog.Info().Msg(fmt.Sprint(msg...))
	}
}

func Debug(msg ...interface{}) {
	if logger.level <= 2 {
		logger.zeroLog.Debug().Msg(fmt.Sprint(msg...))
	}
}

func Error(msg ...interface{}) {
	if logger.level <= 3 {
		logger.zeroLog.Error().Msg(fmt.Sprint(msg...))
	}
}
