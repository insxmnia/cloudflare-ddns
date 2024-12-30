package slogger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func Info(title string, args ...any) {
	slog.Info(title, args...)
	logger.Info(title, args...)
}

func Warn(title string, args ...any) {
	slog.Warn(title, args...)
	logger.Warn(title, args...)
}

func Error(title string, args ...any) {
	slog.Error(title, args...)
	logger.Error(title, args...)
}
func Fatal(title string, args ...any) {
	slog.Error(title, args...)
	logger.Error(title, args...)
	os.Exit(1)
}

func init() {
	once.Do(func() {
		file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		jsonHandler := slog.NewJSONHandler(file, nil)

		logger = slog.New(jsonHandler)

	})
}
