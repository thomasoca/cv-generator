package image

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func mimeCheck(response http.Response) (string, error) {
	contentType := strings.Split(response.Header.Get("Content-Type"), ";")[0]
	extension := strings.Split(contentType, "/")[1]
	if stringInSlice(extension, []string{"jpeg", "jpg", "png"}) {
		return extension, nil
	}
	return "", errors.New("the requested file does not have a proper image extension, please use .jpg or .png only")
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func ImageFromUrl(URL, dirName string) (string, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("received non 200 response code")
	}
	//Check MIME type
	fileExt, err := mimeCheck(*response)
	if err != nil {
		return "", err
	}
	//Create a empty file
	imageName := fmt.Sprintf("image.%s", fileExt)
	fileName := filepath.Join(dirName, imageName)
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

func imageWriter(imageData image.Image, dirName string, extension string) (string, error) {

	//Encode from image format to writer
	imageName := filepath.Join(dirName, "image."+extension)
	f, err := os.OpenFile(imageName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(f, imageData, &jpeg.Options{Quality: 75})
	if err != nil {
		return "", err
	}
	return imageName, err
}

func ImageFromBase64(data string, dirName string) (string, error) {
	coI := strings.Index(data, ",")
	rawImage := data[coI+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(rawImage))

	switch strings.TrimSuffix(data[5:coI], ";base64") {
	case "image/png":
		pngImageData, err := png.Decode(reader)
		if err != nil {
			return "", err
		}
		imageName, err := imageWriter(pngImageData, dirName, "png")
		return imageName, err
	case "image/jpeg":
		jpegImageData, err := jpeg.Decode(reader)
		if err != nil {
			return "", err
		}
		imageName, err := imageWriter(jpegImageData, dirName, "jpeg")
		return imageName, err
	}
	return "", errors.New("unrecognized format")
}
