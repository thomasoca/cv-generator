package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func JsonInput(fname string) []byte {
	jsonFile, err := os.Open(fname)
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func RemoveFiles(dirName string) {
	e := os.RemoveAll(dirName)
	if e != nil {
		panic(e)
	}
}
