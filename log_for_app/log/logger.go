package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"log"
)

type level int

var debugLevel level = 1
var infoLevel level = 2
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

func ZeroLogger() *zerolog.Logger {
	return &logger.zeroLog
}

func DebugMsg(msg ...interface{}) {
	if logger.level <= 1 {
		logger.zeroLog.Debug().Msg(fmt.Sprint(msg...))
	}
}

func Print(msg ...interface{}) {
	if logger.level <= 2 {
		logger.zeroLog.Info().Msg(fmt.Sprint(msg...))
	}
}

func Debug() *zerolog.Event {
	return logger.zeroLog.Debug()
}

func Error() *zerolog.Event {
	return logger.zeroLog.Error()
}

func ErrorMsg(msg ...interface{}) {
	if logger.level <= 3 {
		logger.zeroLog.Error().Msg(fmt.Sprint(msg...))
	}
}

func Println(msg ...interface{}) {
	logger.zeroLog.Print(fmt.Sprint(msg...), "\n")
}

func Printf(format string, msg ...interface{}) {
	logger.zeroLog.Printf(format, msg)
}

func DebugMode() {
	logger.level = debugLevel
}

func InfoMode() {
	logger.level = infoLevel
}

func ErrorMode() {
	logger.level = errorLevel
}
