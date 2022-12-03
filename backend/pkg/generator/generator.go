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

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func CreateFile(user models.User, fileType string) (string, error) {
	var fileName, dirName string
	rand.Seed(time.Now().UnixNano())
	envMode := ""
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	switch envMode {
	case "server":
		// Create the file on tempdir (for prd)
		fileName = randSeq(10)
		randomTempDirName := randSeq(15)
		dirName, err = ioutil.TempDir("", randomTempDirName)
		if err != nil {
			return "", err
		}
	default:
		err := os.Mkdir("test", 0755)
		if err != nil {
			log.Println(err)
		}
		fileName = "test"
		dirName = path + "/test"
	}
	latexFilePath, err := createLatexFile(user, fileName, dirName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if fileType == "pdf" {
		pdfFileName := filepath.Join(dirName, fileName+".pdf")
		err = createResumeFile(dirName, fileName)
		if err != nil {
			return "", err
		}
		return pdfFileName, nil
	}
	return latexFilePath, nil
}
