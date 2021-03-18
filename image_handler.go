package main

import (
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func imageFromUrl(URL, dirName string) (string, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("Received non 200 response code")
	}
	//Create a empty file
	fileName := filepath.Join(dirName, "image.jpg")
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return fileName, err
}

func imageFromBase64(data string, dirName string) (string, error) {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	//Encode from image format to writer
	imageName := filepath.Join(dirName, "image.jpg")
	f, err := os.OpenFile(imageName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)

	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)

	}
	return imageName, err
}
