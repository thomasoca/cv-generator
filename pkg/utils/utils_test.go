package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestJsonInput(t *testing.T) {
	// Create a temporary JSON file for testing
	content := []byte(`{"name": "John", "age": 30}`)
	tmpfile, err := ioutil.TempFile("", "test.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	// Test JsonInput function
	data := JsonInput(tmpfile.Name())
	if !bytes.Equal(data, content) {
		t.Errorf("JsonInput() = %s; want %s", data, content)
	}
}

func TestRemoveFiles(t *testing.T) {
	// Create a temporary directory and a test file inside it
	tmpDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "testfile.txt")
	if err := ioutil.WriteFile(tmpFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}

	// Test RemoveFiles function
	RemoveFiles(tmpFile)

	// Check if the directory has been removed
	_, err = os.Stat(tmpDir)
	if !os.IsNotExist(err) {
		t.Errorf("RemoveFiles() did not remove the directory %s", tmpDir)
	}
}

func TestRunCommand(t *testing.T) {
	cmd := "echo"
	arg := "Hello, World!"

	// Create buffer to capture the command output
	var stdout, stderr bytes.Buffer

	// Test RunCommand function
	err := RunCommand(cmd, &stdout, &stderr, arg)
	if err != nil {
		t.Fatalf("RunCommand() returned an error: %v", err)
	}

	expectedOutput := arg + "\n"
	if stdout.String() != expectedOutput {
		t.Errorf("RunCommand() output = %s; want %s", stdout.String(), expectedOutput)
	}
}
