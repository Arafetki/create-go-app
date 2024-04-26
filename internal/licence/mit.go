package licence

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/arafetki/create-go-app/assets"
)

func GenerateMIT(path string, author string, year int) error {

	licenceTemplate, err := assets.EmbeddedFiles.ReadFile("templates/mit.go.tmpl")
	if err != nil {
		return err
	}

	tmpl, err := template.New("licence").Parse(string(licenceTemplate))
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(path, "LICENSE"))
	if err != nil {
		return err
	}
	defer file.Close()

	data := LicenseData{
		Author: author,
		Year:   year,
	}

	err = tmpl.Execute(file, data)

	return err
}
