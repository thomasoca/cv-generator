package generator

import (
	"os"
	"strings"
	"text/template"
)

func replaceUnescapedChar(str string) string {
	s := strings.NewReplacer(
		"_", "{\\_}",
		"#", "{\\#}",
		"%", "{\\%}",
		"&", "{\\&}",
		"$", "{\\$}",
		"{", "{\\{}",
		"}", "{\\}}",
		"\\", "{\\textbackslash}",
		"^", "{\\textasciicircum}",
		"~", "{\\textasciitilde}",
	)

	return s.Replace(str)
}

func createLatexFile(fg FileGenerator) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	templatePath := path + "/templates/rover.tmpl"
	tpl, err := template.New("rover.tmpl").Funcs(template.FuncMap{"replaceUnescapedChar": replaceUnescapedChar}).ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Convert image
	err = fg.user.Modify(fg.DirPath, fg.output)
	if err != nil {
		return err
	}

	f, err := os.Create(fg.latexPath)
	if err != nil {
		return err
	}
	// Execute the template to the file.
	err = tpl.Execute(f, fg.user)
	if err != nil {
		return err
	}
	// Close the file when done.
	f.Close()

	return nil
}
