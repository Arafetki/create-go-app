package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

var Logger = slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

func LogError(msg string, args ...any) {
	Logger.Error(msg, args...)
}

func LogInfo(msg string, args ...any) {
	Logger.Info(msg, args...)
}

func LogWarning(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

func LogDebug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}
