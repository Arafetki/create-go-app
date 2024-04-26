package cmd

import (
	"fmt"

	"github.com/arafetki/create-go-app/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the create-go-app version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", version.Get())
	},
}
