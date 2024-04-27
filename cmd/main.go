package main

import (
	"os"
	"runtime/debug"

	"github.com/abenk-oss/scaffold-go-app/cmd/scaffgo"
	"github.com/abenk-oss/scaffold-go-app/internal/logger"
)

func main() {

	err := scaffgo.Execute()
	if err != nil {
		trace := string(debug.Stack())
		logger.LogError(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
