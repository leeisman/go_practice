package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	settingZeroLog()
	log.Debug().Msg("test")
}

func settingZeroLog() {
	if true {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
