package log

import (
	"fmt"
	"log"
)

var logger *ll

type ll struct {
	defaultTag string
	customTags []interface{}
	// 1 info debug error , 2 debug error, 3 error
	level level
}
type level int

var debugLevel level = 1
var infoLevel level = 2
var errorLevel level = 3

func init() {
	logger = &ll{
		defaultTag: "go-log",
		level:      infoLevel,
	}
}

func Print(msg ...interface{}) {
	if logger.level <= 1 {
		logger.Info(msg...)
	}
}

func Debug(msg ...interface{}) {
	if logger.level <= 2 {
		logger.Debug(msg...)
	}
}

func Error(msg ...interface{}) {
	if logger.level <= 3 {
		logger.Error(msg...)
	}
}

func (l *ll) Info(msg ...interface{}) {
	log.Print(fmt.Sprintf(fmt.Sprint(l.defaultTag, "[info]"), l.customTags...), msg)
}

func (l *ll) Debug(msg ...interface{}) {
	log.Print(fmt.Sprintf(fmt.Sprint(l.defaultTag, "[debug]"), l.customTags...), msg)
}

func (l *ll) Error(msg ...interface{}) {
	log.Print(fmt.Sprintf(fmt.Sprint(l.defaultTag, "[error]"), l.customTags...), msg)
}

func DebugMode() {
	logger.level = debugLevel
}

func ErrorMode() {
	logger.level = errorLevel
}

func InfoMode() {
	logger.level = infoLevel
}
