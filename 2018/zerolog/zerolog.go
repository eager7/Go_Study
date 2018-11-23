package main

import (
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = ""
	log.Debug().Msg("test")
	log.Print("zero")
}
