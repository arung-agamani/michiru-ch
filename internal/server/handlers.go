package server

import (
	"encoding/json"
	"fmt"
	_ "michiru/docs"
	"michiru/internal/db"
	"michiru/internal/repository"
	"michiru/internal/server/handlers"
	"michiru/internal/services"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @Router /api/v1/projects [post]
// @Router /api/v1/projects [get]
// @Router /api/v1/projects [put]
// @Router /api/v1/projects [delete]
// @Router /api/v1/projects/{id} [get]
func RegisterRoutes(router *mux.Router) {
	dbConn, err := db.Connect()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	projectRepo := repository.NewProjectRepository(dbConn)
	projectHandler := handlers.NewProjectHandler(*projectRepo)

	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/discord", SendMessageHandler).Methods("POST")
	apiV1.HandleFunc("/github-webhook", handlers.HandleGithubWebhook)
	apiV1.HandleFunc("/projects", projectHandler.CreateProject).Methods("POST")
	apiV1.HandleFunc("/projects", projectHandler.ListProjects).Methods("GET")
	apiV1.HandleFunc("/projects/{id}", projectHandler.UpdateProject).Methods("PUT")
	apiV1.HandleFunc("/projects/{id}", projectHandler.DeleteProject).Methods("DELETE")
	apiV1.HandleFunc("/projects/{id}", projectHandler.GetProject).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Konnichiwa, sekai")
}

type MessageRequest struct {
	ChannelID string `json:"channel"`
	Message   string `json:"message"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	discordService, err := services.NewDiscordService()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error initializing Discord service: %v", err), http.StatusInternalServerError)
		return
	}
	defer discordService.Close()

	err = discordService.SendMessage(req.ChannelID, req.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending message: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}
