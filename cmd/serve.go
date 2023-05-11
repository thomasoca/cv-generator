package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/thomasoca/cv-generator/pkg/app"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"ser"},
	Short:   "Serve the backend server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		serveHttpServer(port)
	},
}

func serveHttpServer(port string) {

	mux := http.NewServeMux()
	handlers := new(app.HttpHandlers)
	mux.Handle("/", http.FileServer(http.Dir("./build")))
	mux.HandleFunc("/api/v1/generate", handlers.GenerateFileHandler)
	mux.HandleFunc("/api/v1/example", handlers.ExampleFileHandler)
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	if port == "" {
		port = "8170"
		log.Printf("defaulting to port %s", port)
	}
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringP("port", "p", "", "Port for the http server")
}
