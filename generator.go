package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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

func createFile(user User) (string, error) {
	rand.Seed(time.Now().UnixNano())
	path := os.Getenv("PROJECT_DIR")
	if path == "" {
		localPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		path = localPath
	}
	templatePath := path + "/templates/template.txt"
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	// Create the file
	randomFname := randSeq(10)
	dname, err := ioutil.TempDir("", "tempdir")
	filename := filepath.Join(dname, randomFname+".tex")
	pdfname := filepath.Join(dname, randomFname+".pdf")

	// Convert image
	err = user.Modify(dname)
	if err != nil {
		e := os.RemoveAll(dname)
		if e != nil {
			return "", err
		}
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		e := os.RemoveAll(dname)
		if e != nil {
			return "", err
		}
		return "", err
	}
	// Execute the template to the file.
	err = tpl.Execute(f, user)
	if err != nil {
		log.Println(err)
		e := os.RemoveAll(dname)
		if e != nil {
			return "", err
		}
		return "", err
	}
	err = generateLatex(dname, filename)
	if err != nil {
		e := os.RemoveAll(dname)
		if e != nil {
			return "", err
		}
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
	// for debug locally or in docker
	// cmdArgs := []string{"-interaction=nonstopmode", "-synctex=1", "-halt-on-error", filename}

	cmd := exec.Command(app, cmdArgs...)
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return err
	}
	return err
}
