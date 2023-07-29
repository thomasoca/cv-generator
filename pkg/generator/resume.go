package generator

import (
	"bytes"
	"errors"
	"log"

	"github.com/thomasoca/cv-generator/pkg/utils"
)

func createResumeFile(fg FileGenerator) error {
	app := "pdflatex"
	outdir := "-output-directory=" + fg.DirPath
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmdArgs := []string{outdir, "-interaction=nonstopmode", "-synctex=1", "-halt-on-error", fg.latexPath}
	err := utils.RunCommand(app, &stdout, &stderr, cmdArgs...)
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		log.Println(stdout.String())
		return errors.New("there is a problem when running latex in the server")
	}
	log.Println("Latex file generated successfully")
	return err
}
