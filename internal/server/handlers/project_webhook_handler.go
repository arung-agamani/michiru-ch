package handlers

import (
	"encoding/json"
	"log"
	"michiru/internal/models"
	"michiru/internal/repository"
	"michiru/internal/services"
	"michiru/internal/utils"
	"net/http"

	"encoding/base64"

	"github.com/go-playground/webhooks/v6/github"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProjectWebhookHandler struct {
	Repo repository.ProjectRepository
}

func NewProjectWebhookHandler(repo repository.ProjectRepository) *ProjectWebhookHandler {
	return &ProjectWebhookHandler{Repo: repo}
}

type updateWebhook struct {
	WebhookOrigin string `json:"webhook_origin"`
	// WebhookURL    string `json:"webhook_url"`
	WebhookSecret string `json:"webhook_secret"`
}

// UpdateWebhook godoc
// @Summary Update a project's webhook details by ID
// @Description Updates a project's webhook details using the provided ID and request body
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Project ID"
// @Param project body updateWebhook true "Updated project data"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{id}/webhook [put]
func (h *ProjectWebhookHandler) UpdateWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
		return
	}

	var updateWebhook updateWebhook
	if err := json.NewDecoder(r.Body).Decode(&updateWebhook); err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve project"})
		return
	}

	project.ID = id
	project.WebhookOrigin = updateWebhook.WebhookOrigin
	// project.WebhookURL = updateWebhook.WebhookURL
	project.WebhookSecret = updateWebhook.WebhookSecret

	if err := h.Repo.UpdateWebhook(project); err != nil {
		log.Printf("Error updating project webhook: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to update project webhook"})
		return
	}

	utils.WriteSuccessJSON(w, project)
}

// GenerateWebhook godoc
// @Summary Generate a new webhook for a project by ID
// @Description Generates a new webhook for a project using the provided ID
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Project ID"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{id}/webhook [post]
func (h *ProjectWebhookHandler) GenerateWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve project"})
		return
	}

	project.ID = id
	project.WebhookOrigin = "github"

	uuidValue := uuid.New()
	encodedUUID := base64.StdEncoding.EncodeToString([]byte(uuidValue.String()))
	project.WebhookURL = encodedUUID
	project.WebhookSecret = ""

	if err := h.Repo.UpdateWebhook(project); err != nil {
		log.Printf("Error updating project webhook: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to update project webhook"})
		return
	}

	utils.WriteSuccessJSON(w, project)
}

// HandleWebhookPayload godoc
// @Summary Handle webhook payload
// @Description Handles the webhook payload for a project using the provided ID and webhook URL
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Project ID"
// @Param webhookUrl path string true "Webhook URL"
// @Success 200 {object} utils.Response{data=string}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{id}/webhook/{webhookUrl} [post]
func (h *ProjectWebhookHandler) HandleWebhookPayload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		log.Printf("Missing project ID")
		utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve project"})
		return
	}

	webhookUrl := vars["webhookId"]
	if webhookUrl == "" {
		log.Printf("Missing webhook URL")
		utils.WriteBadRequestJSON(w, []string{"Missing webhook URL"})
		return
	}
	if project.WebhookURL != webhookUrl {
		log.Printf("Invalid webhook URL")
		utils.WriteBadRequestJSON(w, []string{"Invalid webhook URL"})
		return
	}

	hook, _ := github.New(github.Options.Secret(project.WebhookSecret))
	payload, err := hook.Parse(r, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Printf("Event not found")
		} else {
			log.Printf("Error parsing payload: %v", err)
		}
		utils.WriteBadRequestJSON(w, []string{"Invalid payload."})
		return
	}

	switch payload := payload.(type) {

	case github.PushPayload:
		push := payload
		log.Printf("Received push event for %s", push.Repository.FullName)
		sendPushEventNotification(push, project)
		utils.WriteSuccessJSON(w, "Received push event")
	default:
		log.Printf("Received unsupported event")
		utils.WriteBadRequestJSON(w, []string{"Invalid payload. Only accepting PushEvent"})
		return
	}
}

func sendPushEventNotification(payload github.PushPayload, project *models.Project) {
	// Send a notification to the Discord channel
	discordService, err := services.NewDiscordService()
	if err != nil {
		log.Printf("Error initializing Discord service: %v", err)
		return
	}

	message := "New push event on repository " + payload.Repository.FullName
	message += "\nCommit message: " + payload.Commits[0].Message
	message += "\nCommit URL: " + payload.Commits[0].URL
	message += "\nPusher: " + payload.Pusher.Name
	message += "\nBranch: " + payload.Ref

	if err := discordService.SendMessage(project.ChannelID, message); err != nil {
		log.Printf("Error sending message: %v", err)
	}

}
