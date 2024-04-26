package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	cmd "github.com/arafetki/create-go-app/cmd/cli/commands"
	"github.com/lmittmann/tint"
)

func main() {

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	err := cmd.Execute()
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)

	}
}
