package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user User
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
		fname, err := createFile(user)
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

func JsonInput(fname string) []byte {
	jsonFile, err := os.Open(fname)
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func getExample(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-type", "application/json")
		body := JsonInput("./examples/user.json")
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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./build")))
	mux.HandleFunc("/api/generate", serveFile)
	mux.HandleFunc("/api/example", getExample)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}

}
