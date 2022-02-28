package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func replaceUnescapedChar(str string) string {
	return strings.Replace(str, "_", "{\\_}", -1)
}

func (u *User) Modify(dirname string) error {
	imageData := u.PersonalInfo.Picture
	if imageData != "" {
		var newImage string
		checkUrl := IsUrl(imageData)
		if checkUrl {
			newImage, err := imageFromUrl(imageData, dirname)
			if err != nil {
				return err
			}
			u.PersonalInfo.Picture = newImage
			return err
		}

		newImage, err := imageFromBase64(imageData, dirname)
		if err != nil {
			return err
		}
		u.PersonalInfo.Picture = newImage
		return err
	}

	return nil
}

func removeLatexFiles(env string, dirName string) {
	if env == "PRD" {
		e := os.RemoveAll(dirName)
		if e != nil {
			panic(e)
		}
	}
}

func createFile(user User) (string, error) {
	rand.Seed(time.Now().UnixNano())
	path := os.Getenv("PROJECT_DIR")
	if path == "" {
		localPath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = localPath
	}
	templatePath := path + "/templates/template.txt"
	tpl, err := template.New("template.txt").Funcs(template.FuncMap{"replaceUnescapedChar": replaceUnescapedChar}).ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	envMode := os.Getenv("ENV_MODE")
	var fName, dname string
	switch envMode {
	case "PRD":
		// Create the file on tempdir (for prd)
		fName = randSeq(10)
		randomTempDirName := randSeq(15)
		dname, err = ioutil.TempDir("", randomTempDirName)
		if err != nil {
			return "", err
		}
	default:
		err := os.Mkdir("test", 0755)
		if err != nil {
			log.Println(err)
		}
		fName = "test"
		dname = path + "/test"
	}

	filename := filepath.Join(dname, fName+".tex")
	pdfname := filepath.Join(dname, fName+".pdf")

	// Convert image
	err = user.Modify(dname)
	if err != nil {
		removeLatexFiles(envMode, dname)
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		removeLatexFiles(envMode, dname)
		return "", err
	}
	// Execute the template to the file.
	err = tpl.Execute(f, user)
	if err != nil {
		removeLatexFiles(envMode, dname)
		return "", err
	}
	err = generateLatex(dname, filename)
	if err != nil {
		removeLatexFiles(envMode, dname)
		return "", err
	}
	// Close the file when done.
	f.Close()

	return pdfname, nil
}

func generateLatex(dirname string, filename string) error {
	app := "pdflatex"
	outdir := "-output-directory=" + dirname
	cmdArgs := []string{outdir, "-interaction=nonstopmode", "-synctex=1", "-halt-on-error", filename}
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
	return err
}
