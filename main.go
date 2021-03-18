package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad request"}`))
	}
	fname, err := createFile(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Internal server error"}`))
	}

	// Open file
	f, err := os.Open(fname)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()
	// Set header
	w.Header().Set("Content-type", "application/pdf")

	// Stream to response
	if _, err := io.Copy(w, f); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	e := os.RemoveAll(filepath.Dir(fname))
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	http.HandleFunc("/user", serveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
