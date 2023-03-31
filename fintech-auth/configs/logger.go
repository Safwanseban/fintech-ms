package configs

import (
	"os"

	"github.com/rs/zerolog"
)

func Getlogger() *zerolog.Logger {

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Debug().Enabled()
	logger.WithLevel(logger.GetLevel()).Enabled()
	return &logger

}
