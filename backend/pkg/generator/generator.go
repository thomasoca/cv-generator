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

func (f *FileGenerator) PathGenerator(user models.User, output string) error {
	f.user = user
	rand.Seed(time.Now().UnixNano())
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	switch output {
	case "server":
		// Create the file on tempdir (for prd)
		f.fileName = randSeq(10)
		randomTempDirName := randSeq(15)
		f.DirPath, err = ioutil.TempDir("", randomTempDirName)
		if err != nil {
			return err
		}
	case "app":
		_ = os.Mkdir("result", 0755)
		f.fileName = user.PersonalInfo.Name + " Resume"
		f.DirPath = path + "/result"
	case "development":
		_ = os.Mkdir("test", 0755)
		f.fileName = "test"
		f.DirPath = path + "/test"
		log.Println("testing output created at ", f.DirPath)
	default:
		p := output + "/result"
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		f.fileName = user.PersonalInfo.Name + " Resume"
		f.DirPath = p
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

func CreateFile(user models.User, output string) (string, error) {
	var generator FileGenerator

	err := generator.PathGenerator(user, output)
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
