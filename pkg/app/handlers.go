package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/thomasoca/cv-generator/pkg/generator"
	"github.com/thomasoca/cv-generator/pkg/models"
	"github.com/thomasoca/cv-generator/pkg/utils"
)

type HttpHandlers struct {
	DevMode bool
}

func (h *HttpHandlers) GenerateFileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user models.User
		err := decoder.Decode(&user)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(`{"message": "Bad request"}`))
			if err != nil {
				log.Panic(err)
			}
			return
		}
		fname, err := generator.CreateFile(user, "server")
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			if fname != "" {
				utils.RemoveFiles(fname)
			}
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
			utils.RemoveFiles(fname)
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
			utils.RemoveFiles(fname)
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`{"message": "Failed sending file"}`))
			if err != nil {
				log.Panic(err)
			}
			return
		}

		utils.RemoveFiles(fname)
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

func (h *HttpHandlers) ExampleFileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-type", "application/json")
		body := utils.JsonInput("./examples/user.json")
		_, err := w.Write(body)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`{"message": "Failed processing file"}`))
			if err != nil {
				log.Panic(err)
			}
		}
		return
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
