package app

import (
	"log"
	"net/http"

	"github.com/thomasoca/cv-generator/backend/pkg/utils"
)

func ExampleFileHandler(w http.ResponseWriter, r *http.Request) {
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
