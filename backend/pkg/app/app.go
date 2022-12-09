package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/thomasoca/cv-generator/backend/pkg/generator"
	"github.com/thomasoca/cv-generator/backend/pkg/models"
)

type AppCmd struct {
	InputPath  string
	OutputPath string
}

func (a *AppCmd) GenerateFile() (string, error) {
	var user models.User
	j, err := os.Open(a.InputPath)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer j.Close()
	b, _ := ioutil.ReadAll(j)
	json.Unmarshal(b, &user)
	f, err := generator.CreateFile(user, a.OutputPath)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return f, nil

}
