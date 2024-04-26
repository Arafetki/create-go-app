package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "create-go-app",
		Short: "A generator for Golang based applications",
		Long:  "This application is a tool to effortlessly generate the ideal application scaffold for your Golang projects",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	rootCmd.AddCommand(versionCmd)
}
