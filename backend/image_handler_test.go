package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestImageFromUrl(t *testing.T) {
	localPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	URL := "https://trimelive.com/wp-content/uploads/2020/12/Gambar-wa-12.jpg"
	err = os.Mkdir("test_image", 0755)
	if err != nil {
		log.Println(err)
	}
	dname := localPath + "/test_image"
	fname, err := imageFromUrl(URL, dname)
	if err != nil {
		log.Println(err)
		t.Errorf("image fetching process was failed")
	}
	expectedFname := localPath + "/test_image/image.jpeg"

	// Check filename result
	if fname != expectedFname {
		t.Errorf("image fetching was incorrect, got: %s, want: %s.", fname, expectedFname)
	}

	// Check file exsistence
	_, err = os.Stat(fname)
	if os.IsNotExist(err) {
		t.Errorf("failed to create image file")
	}
	err = os.RemoveAll(filepath.Dir(fname))
	if err != nil {
		log.Panic(err)
	}
}
