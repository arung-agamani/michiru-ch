package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"michiru/internal/services"
	"michiru/internal/utils"
)

type GithubWebhookPayload struct {
	Ref string `json:"ref"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
}

func HandleGithubWebhook(w http.ResponseWriter, r *http.Request) {
	var payload GithubWebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Repository: %s\nBranch: %s\nPushed by: %s", 
		payload.Repository.FullName, payload.Ref, payload.Pusher.Name)
	
	discordService, err := services.NewDiscordService()
	if err != nil {
		utils.WriteBadRequestJSON(w, []string{err.Error()})
		return
	}

	err = discordService.SendMessage("DISCORD_CHANNEL_ID_HERE", message)
	if err != nil {
		utils.WriteInternalServerErrorJSON(w, []string{err.Error()})
		return
	}

	utils.WriteSuccessJSON(w, "webhook processed successfully")
}