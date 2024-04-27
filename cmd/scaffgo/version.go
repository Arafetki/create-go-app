package scaffgo

import (
	"fmt"

	"github.com/abenk-oss/scaffold-go-app/internal/version"
	"github.com/spf13/cobra"
)

var (
	v *version.Info
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the scaffgo version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", v.CLIVersion)
		fmt.Printf("Go version: %s\n", v.RuntimeVersion)
		fmt.Printf("Git commit: %s\n", v.CommitHash)
	},
}

func init() {
	v = version.New("0.1.0")
	rootCmd.AddCommand(versionCmd)
}
