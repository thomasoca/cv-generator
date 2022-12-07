package generator

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/thomasoca/cv-generator/backend/pkg/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type FileGenerator struct {
	user      models.User
	latexPath string
	pdfPath   string
	DirPath   string
	fileName  string
}

func (f *FileGenerator) PathGenerator(user models.User, devMode bool) error {
	f.user = user
	rand.Seed(time.Now().UnixNano())
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	if !devMode {
		// Create the file on tempdir (for prd)
		f.fileName = randSeq(10)
		randomTempDirName := randSeq(15)
		f.DirPath, err = ioutil.TempDir("", randomTempDirName)
		if err != nil {
			return err
		}
	} else {
		err := os.Mkdir("test", 0755)
		if err != nil {
			log.Println(err)
			return err
		}
		f.fileName = "test"
		f.DirPath = path + "/test"
		log.Println("testing output created in ", f.DirPath)
	}
	f.latexPath = filepath.Join(f.DirPath, f.fileName+".tex")
	f.pdfPath = filepath.Join(f.DirPath, f.fileName+".pdf")
	return nil
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateFile(user models.User, devMode bool) (string, error) {
	var generator FileGenerator

	err := generator.PathGenerator(user, devMode)
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = createLatexFile(generator)
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = createResumeFile(generator)
	if err != nil {
		return "", err
	}
	return generator.pdfPath, nil
}
