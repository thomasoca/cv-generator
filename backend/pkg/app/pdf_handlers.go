package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/thomasoca/cv-generator/backend/pkg/generator"
	"github.com/thomasoca/cv-generator/backend/pkg/models"
)

func GenerateFileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user models.User
		err := decoder.Decode(&user)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(`{"message": "Bad request"}`))
			if err != nil {
				log.Panic(err)
			}
			return
		}
		fname, err := generator.CreateFile(user)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`{"message": "Failed creating file"}`))
			if err != nil {
				log.Panic(err)
			}
			return
		}

		// Open file
		f, err := os.Open(fname)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`{"message": "Failed processing file"}`))
			if err != nil {
				log.Panic(err)
			}
			return
		}
		defer f.Close()
		// Set header
		w.Header().Set("Content-type", "application/pdf")

		// Stream to response
		if _, err := io.Copy(w, f); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`{"message": "Failed sending file"}`))
			if err != nil {
				log.Panic(err)
			}
			return
		}

		err = os.RemoveAll(filepath.Dir(fname))
		if err != nil {
			log.Panic(err)
		}
	default:
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(`{"message": "Method not allowed"}`))
		if err != nil {
			log.Panic(err)
		}
		return
	}
}