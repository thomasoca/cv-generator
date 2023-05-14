package image

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
	URL := "https://picsum.photos/200/300"
	err = os.Mkdir("test_image", 0755)
	if err != nil {
		log.Println(err)
	}
	dname := localPath + "/test_image"
	fname, err := ImageFromUrl(URL, dname)
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

func TestIsUrl(t *testing.T) {
	url := "https://localhost:8170/test.jpg"
	check := IsUrl(url)
	if !check {
		t.Errorf("URL checking was incorrect, got: false, want: true")
	}
}

func TestIsFile(t *testing.T) {
	path1 := "/home/thomasoca/Documents/important/foto.jpg"
	check := IsImageFileExist(path1)
	if !check {
		t.Errorf("directory checking was incorrect, got: true, want: false")
	}
}
