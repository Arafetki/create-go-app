package cmd

import (
	"fmt"
	"runtime"

	"github.com/arafetki/create-go-app/internal/version"
	"github.com/spf13/cobra"
)

var (
	GoVersion  string
	Version    = "unknown-version"
	GitCommit  = "unknown-commit"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display the create-go-app version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Go version: %s\n", GoVersion)
			fmt.Printf("Git commit: %s\n", GitCommit)
		},
	}
)

func init() {
	Version = version.Get()
	GoVersion = runtime.Version()
	GitCommit = version.GetCommit()

	rootCmd.AddCommand(versionCmd)

}
