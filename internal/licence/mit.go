package licence

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/arafetki/create-go-app/assets"
)

func GenerateMIT(path string, copyright Copyright) error {

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

	err = tmpl.Execute(file, copyright)

	return err
}
