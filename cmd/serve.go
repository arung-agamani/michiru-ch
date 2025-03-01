package cmd

import (
	"fmt"
	"log"
	"michiru/internal/server"
	"net/http"

	"github.com/spf13/cobra"
)

var port string

var serveCmd = &cobra.Command{
	Use: "server",
	Short: "Starts a web server",
	Long: "This command starts a web server handling Michiru-ch specs",
	Run: serveHandler,
}

func serveHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Starting web server on port", port)
	router := server.NewRouter()
	
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the web server")
}