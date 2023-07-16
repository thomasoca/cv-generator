package utils

import (
	"bytes"
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

func RunCommand(command string, stdout, stderr *bytes.Buffer, args ...string) error {
	cmd := exec.Command(command, args...)
	if stdout != nil {
		cmd.Stdout = stdout
	} else {
		cmd.Stdout = os.Stdout
	}
	if stderr != nil {
		cmd.Stderr = stderr
	} else {
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}
