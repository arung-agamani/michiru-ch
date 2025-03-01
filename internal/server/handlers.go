package server

import (
	"encoding/json"
	"fmt"
	"michiru/internal/server/handlers"
	"michiru/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/discord", SendMessageHandler).Methods("POST")
	router.HandleFunc("/github-webhook", handlers.HandleGithubWebhook)
	router.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Konnichiwa, sekai")
}

type MessageRequest struct {
	ChannelID string `json:"channel"`
	Message string `json:"message"`
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