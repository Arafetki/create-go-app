package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arafetki/create-go-app/internal/licence"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Run: func(cmd *cobra.Command, args []string) {

		var projectName string

		if len(args) == 0 {
			projectName = StringPrompt("Enter Project name")

		} else {
			projectName = args[0]
		}

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}

		projectPath := filepath.Join(currentDir, projectName)

		if _, err := os.Stat(projectPath); os.IsNotExist(err) {

			err = os.Mkdir(projectName, 0755)
			if err != nil {
				fmt.Println("Error creating project directory:", err.Error())
				return
			}
		} else {
			if !YesNoPrompt("Do you want to override existing files?", false) {
				fmt.Println("Aborting initialization.")
				return
			} else {
				fmt.Println(("Overriding existing files..."))
			}
		}

		err = licence.GenerateMIT(filepath.Join(currentDir, projectName), "Arafet BenKilani", 2024)
		if err != nil {
			fmt.Println("Error generating LICENSE file:", err.Error())
			return
		}

		fmt.Println("Project initialized successfully.")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
