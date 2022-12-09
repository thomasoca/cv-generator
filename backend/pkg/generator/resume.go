package generator

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

func createResumeFile(fg FileGenerator) error {
	app := "pdflatex"
	outdir := "-output-directory=" + fg.DirPath
	cmdArgs := []string{outdir, "-interaction=nonstopmode", "-synctex=1", "-halt-on-error", fg.latexPath}
	var stderr bytes.Buffer
	var out bytes.Buffer
	cmd := exec.Command(app, cmdArgs...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		log.Println(out.String())
		return errors.New("there is a problem when running latex in the server")
	}
	log.Println("Latex file generated successfully")
	return err
}
