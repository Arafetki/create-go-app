package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Run: func(cmd *cobra.Command, args []string) {

		// @Todo

		fmt.Println("Do something!")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
