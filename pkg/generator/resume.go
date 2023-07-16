package generator

import (
	"errors"
	"log"

	"github.com/thomasoca/cv-generator/pkg/utils"
)

func createResumeFile(fg FileGenerator) error {
	app := "pdflatex"
	outdir := "-output-directory=" + fg.DirPath
	cmdArgs := []string{outdir, "-interaction=nonstopmode", "-synctex=1", "-halt-on-error", fg.latexPath}
	err := utils.RunCommand(app, cmdArgs...)
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return errors.New("there is a problem when running latex in the server")
	}
	log.Println("Latex file generated successfully")
	return err
}

func CheckVersion() error {
	err := utils.RunCommand("pdflatex", "-version")
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return errors.New("latex backend not available")
	}
	return err
}
