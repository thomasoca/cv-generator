package utils

import (
	"io/ioutil"
	"log"
	"os"
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
