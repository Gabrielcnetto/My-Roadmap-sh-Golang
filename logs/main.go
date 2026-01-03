package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	log.Info().Msg("Info message")
	log.Error().Msg("Error message")
}
