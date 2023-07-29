package image

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
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

func TestImageFromBase64(t *testing.T) {
	// Test a valid base64-encoded PNG image
	base64Data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAApgAAAKYB3X3/OAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAANCSURBVEiJtZZPbBtFFMZ/M7ubXdtdb1xSFyeilBapySVU8h8OoFaooFSqiihIVIpQBKci6KEg9Q6H9kovIHoCIVQJJCKE1ENFjnAgcaSGC6rEnxBwA04Tx43t2FnvDAfjkNibxgHxnWb2e/u992bee7tCa00YFsffekFY+nUzFtjW0LrvjRXrCDIAaPLlW0nHL0SsZtVoaF98mLrx3pdhOqLtYPHChahZcYYO7KvPFxvRl5XPp1sN3adWiD1ZAqD6XYK1b/dvE5IWryTt2udLFedwc1+9kLp+vbbpoDh+6TklxBeAi9TL0taeWpdmZzQDry0AcO+jQ12RyohqqoYoo8RDwJrU+qXkjWtfi8Xxt58BdQuwQs9qC/afLwCw8tnQbqYAPsgxE1S6F3EAIXux2oQFKm0ihMsOF71dHYx+f3NND68ghCu1YIoePPQN1pGRABkJ6Bus96CutRZMydTl+TvuiRW1m3n0eDl0vRPcEysqdXn+jsQPsrHMquGeXEaY4Yk4wxWcY5V/9scqOMOVUFthatyTy8QyqwZ+kDURKoMWxNKr2EeqVKcTNOajqKoBgOE28U4tdQl5p5bwCw7BWquaZSzAPlwjlithJtp3pTImSqQRrb2Z8PHGigD4RZuNX6JYj6wj7O4TFLbCO/Mn/m8R+h6rYSUb3ekokRY6f/YukArN979jcW+V/S8g0eT/N3VN3kTqWbQ428m9/8k0P/1aIhF36PccEl6EhOcAUCrXKZXXWS3XKd2vc/TRBG9O5ELC17MmWubD2nKhUKZa26Ba2+D3P+4/MNCFwg59oWVeYhkzgN/JDR8deKBoD7Y+ljEjGZ0sosXVTvbc6RHirr2reNy1OXd6pJsQ+gqjk8VWFYmHrwBzW/n+uMPFiRwHB2I7ih8ciHFxIkd/3Omk5tCDV1t+2nNu5sxxpDFNx+huNhVT3/zMDz8usXC3ddaHBj1GHj/As08fwTS7Kt1HBTmyN29vdwAw+/wbwLVOJ3uAD1wi/dUH7Qei66PfyuRj4Ik9is+hglfbkbfR3cnZm7chlUWLdwmprtCohX4HUtlOcQjLYCu+fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII="
	dirName := os.TempDir()
	imageName, err := ImageFromBase64(base64Data, dirName)
	if err != nil {
		t.Errorf("Error creating image from base64 data: %s", err)
	}
	defer os.Remove(imageName)
}

func TestIsUrl(t *testing.T) {
	// Test a valid URL
	validURL := "https://example.com/image.jpg"
	if !IsUrl(validURL) {
		t.Errorf("Expected IsUrl to return true for a valid URL, but it returned false.")
	}

	// Test an invalid URL
	invalidURL := "not_a_url"
	if IsUrl(invalidURL) {
		t.Errorf("Expected IsUrl to return false for an invalid URL, but it returned true.")
	}
}

func TestIsImageFileExist(t *testing.T) {
	// Test an existing image file
	// Get a temporary directory for the random file
	tempDir := os.TempDir()

	// Generate a random file name using Unix time
	unixTime := time.Now().Unix()
	randomFileName := fmt.Sprintf("testfile_%d.txt", unixTime)

	// Create the file path
	existingFile := filepath.Join(tempDir, randomFileName)

	// Create an empty file
	file, err := os.Create(existingFile)
	if err != nil {
		log.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	if !IsImageFileExist(existingFile) {
		t.Errorf("Expected IsImageFileExist to return true for an existing file, but it returned false.")
	}

	defer os.Remove(existingFile)

	// Test a non-existing file
	nonExistingFile := "examples/non_existing.jpg"
	if IsImageFileExist(nonExistingFile) {
		t.Errorf("Expected IsImageFileExist to return false for a non-existing image file, but it returned true.")
	}
}
