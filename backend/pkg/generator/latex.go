package generator

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/thomasoca/cv-generator/backend/pkg/models"
	"github.com/thomasoca/cv-generator/backend/pkg/utils"
)

func replaceUnescapedChar(str string) string {
	return strings.ReplaceAll(str, "_", "{\\_}")
}

func createLatexFile(user models.User, fileName string, dirName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	templatePath := path + "/templates/template.tmpl"
	tpl, err := template.New("template.tmpl").Funcs(template.FuncMap{"replaceUnescapedChar": replaceUnescapedChar}).ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	filename := filepath.Join(dirName, fileName+".tex")

	// Convert image
	err = user.Modify(dirName)
	if err != nil {
		utils.RemoveFiles(dirName)
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		utils.RemoveFiles(dirName)
		return "", err
	}
	// Execute the template to the file.
	err = tpl.Execute(f, user)
	if err != nil {
		utils.RemoveFiles(dirName)
		return "", err
	}
	// Close the file when done.
	f.Close()

	return filename, nil
}
