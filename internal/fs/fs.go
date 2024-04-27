package fs

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"

	"github.com/abenk-oss/scaffold-go-app/assets"
)

func GenerateFileFromTemplate(fileName string, templateFilePath string, variables map[string]any) error {

	if fileName == "" {
		return errors.New("not correct or empty file name")
	}

	cleanFileName := filepath.Clean(fileName)

	templateFile, err := assets.EmbeddedFiles.ReadFile(templateFilePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(cleanFileName).Parse(string(templateFile))
	if err != nil {
		return err
	}

	file, err := os.Create(cleanFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	if err = tmpl.Execute(file, variables); err != nil {
		return err
	}

	return nil

}

func RemoveFolders(rootFolder string, folders []string) {
	for _, folder := range folders {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
