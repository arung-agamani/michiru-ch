package handlers

import (
	"encoding/json"
	"log"
	"michiru/internal/models"
	"michiru/internal/repository"
	"michiru/internal/services"
	"michiru/internal/utils"
	"net/http"

	"bytes"
	"encoding/base64"
	"text/template"

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
	project.WebhookOrigin = &updateWebhook.WebhookOrigin
	// project.WebhookURL = updateWebhook.WebhookURL
	project.WebhookSecret = &updateWebhook.WebhookSecret

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
	origin := "github"
	project.WebhookOrigin = &origin

	uuidValue := uuid.New()
	encodedUUID := base64.StdEncoding.EncodeToString([]byte(uuidValue.String()))
	project.WebhookURL = &encodedUUID
	emptySecret := ""
	project.WebhookSecret = &emptySecret

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
	if project.WebhookURL == nil || *project.WebhookURL != webhookUrl {
		log.Printf("Invalid webhook URL")
		utils.WriteBadRequestJSON(w, []string{"Invalid webhook URL"})
		return
	}

	hook, _ := github.New(github.Options.Secret(*project.WebhookSecret))
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

		templateRepo := repository.NewTemplateRepository(h.Repo.DB)
		templates, err := templateRepo.GetByProjectID(id)
		if err != nil {
			log.Printf("Error getting templates: %v", err)
			utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve templates"})
			return
		}

		var templateContent string
		for _, template := range templates {
			if template.EventType == "push" {
				templateContent = template.Template
				break
			}
		}

		if templateContent == "" {
			log.Printf("No template found for push event")
			utils.WriteBadRequestJSON(w, []string{"No template found for push event"})
			return
		}

		message := formatTemplate(templateContent, payload)

		sendPushEventNotification(message, project)
		utils.WriteSuccessJSON(w, "Received push event")
	default:
		log.Printf("Received unsupported event")
		utils.WriteBadRequestJSON(w, []string{"Invalid payload. Only accepting PushEvent"})
		return
	}
}

func formatTemplate(templateContent string, data any) string {
	// Use Go's text/template package to format the template
	tmpl, err := template.New("webhookTemplate").Parse(templateContent)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return "Error formatting template"
	}

	var formattedMessage bytes.Buffer
	if err := tmpl.Execute(&formattedMessage, data); err != nil {
		log.Printf("Error executing template: %v", err)
		return "Error formatting template"
	}

	return formattedMessage.String()
}

func sendPushEventNotification(message string, project *models.Project) {
	// Send a notification to the Discord channel
	discordService, err := services.NewDiscordService()
	if err != nil {
		log.Printf("Error initializing Discord service: %v", err)
		return
	}

	if err := discordService.SendMessage(project.ChannelID, message); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
