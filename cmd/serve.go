package cmd

import (
	"context"
	"fmt"
	"log"
	"michiru/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exited properly")
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the web server")
}