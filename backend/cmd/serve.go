package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/thomasoca/cv-generator/backend/pkg/app"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"ser"},
	Short:   "Serve the backend server",
	Run: func(cmd *cobra.Command, args []string) {
		serveHttpServer()
	},
}

func serveHttpServer() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./build")))
	mux.HandleFunc("/api/v1/generate", app.GenerateFileHandler)
	mux.HandleFunc("/api/v1/example", app.ExampleFileHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
