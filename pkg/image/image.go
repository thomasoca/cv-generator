package image

import (
	"bytes"
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

func IsImageFileExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
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

func ImageFromBase64(base64String string, dirName string) (string, error) {
	// Convert base64 string into byte
	trimmedBase64String := strings.TrimPrefix(base64String, "data:image/jpeg;base64,")
	trimmedBase64String = strings.TrimPrefix(trimmedBase64String, "data:image/png;base64,")
	imageData, err := base64.StdEncoding.DecodeString(trimmedBase64String)
	if err != nil {
		return "", err
	}

	// Get image extension
	parts := strings.Split(base64String, ",")
	metadata := parts[0]
	typeAndEncoding := strings.Split(metadata, ":")
	fileType := typeAndEncoding[1]
	base64Extension := strings.Split(fileType, "/")[1]
	extension := strings.Split(base64Extension, ";")[0]

	// Create image file
	imageName := filepath.Join(dirName, "image."+extension)
	f, err := os.Create(imageName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Create an image.Image from the decoded data
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return "", err
	}

	// Save the image to the file with the appropriate format
	switch extension {
	case "jpeg":
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: 75})
	case "png":
		err = png.Encode(f, img)
	default:
		err = fmt.Errorf("unsupported image format: %s", extension)
	}
	if err != nil {
		return "", err
	}
	return imageName, nil
}
