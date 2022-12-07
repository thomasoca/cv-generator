package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/thomasoca/cv-generator/backend/pkg/app"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"ser"},
	Short:   "Serve the backend server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		mode, _ := cmd.Flags().GetBool("development")
		serveHttpServer(port, mode)
	},
}

func serveHttpServer(port string, mode bool) {

	mux := http.NewServeMux()
	handlers := app.HttpHandlers{DevMode: mode}
	mux.Handle("/", http.FileServer(http.Dir("./build")))
	mux.HandleFunc("/api/v1/generate", handlers.GenerateFileHandler)
	mux.HandleFunc("/api/v1/example", handlers.ExampleFileHandler)
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	if mode {
		log.Println("running server in development mode")
	}
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().String("port", "", "Port for the http server")
	rootCmd.PersistentFlags().Bool("development", false, "Starting the http server in development mode")
}
