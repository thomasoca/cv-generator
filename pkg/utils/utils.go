package utils

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func JsonInput(fname string) []byte {
	jsonFile, err := os.Open(fname)
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func RemoveFiles(fileName string) {
	e := os.RemoveAll(filepath.Dir(fileName))
	if e != nil {
		panic(e)
	}
}

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
