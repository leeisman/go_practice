package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	dlog "log"
)

func main() {
	settingZeroLog()
	log.Logger = zerolog.New(dlog.Writer()).With().Timestamp().Str("tag", "go-log").Logger()
	dlog.Println("go-log", "log says 3")
	log.Debug().Msg("test")
	dlog.Print("test")
}

func settingZeroLog() {
	if true {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
