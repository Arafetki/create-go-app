package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/abenk-oss/scaffold-go-app/cmd/scaffgo"
)

func main() {

	err := scaffgo.Execute()
	if err != nil {
		trace := string(debug.Stack())
		slog.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
