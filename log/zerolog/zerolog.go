package zerolog

import (
	"github.com/rs/zerolog"
)

type Logger struct {
	log *zerolog.Logger
}

func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}
