package xlogger

import (
	"io"
	"os"
	"path/filepath"
	"simple-withdraw-api/internal/config"
	"time"

	"github.com/rs/zerolog"
)

var (
	Logger *zerolog.Logger
)

func Setup(cfg config.Config) {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			panic("Failed to create logs directory: " + err.Error())
		}
	}

	logFile := filepath.Join(logDir, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	var writers []io.Writer

	if cfg.IsDevelopment {
		l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		l.Level(zerolog.DebugLevel)
		Logger = &l
		return
	}
	writers = append(writers, os.Stderr, file)
	multi := io.MultiWriter(writers...)

	l := zerolog.New(multi).With().Timestamp().Logger()
	Logger = &l
}
