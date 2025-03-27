package handlers

import (
	"encoding/json"
	"log"
	"michiru/internal/models"
	"michiru/internal/repository"
	"michiru/internal/services"
	"michiru/internal/utils"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgconn"
)

type ProjectHandler struct {
	Repo repository.ProjectRepository
}

func NewProjectHandler(repo repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{Repo: repo}
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project from the request body
// @Tags projects
// @Accept  json
// @Produce  json
// @Param project body models.Project true "Project to create"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects [post]
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	project.ID = uuid.New().String()
	project.CreatedAt = time.Now().Format(time.RFC3339)
	project.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := h.Repo.Insert(&project); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			utils.WriteBadRequestJSON(w, []string{"Project with the same name already exists"})
			return
		}
		log.Printf("Error inserting project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to create a project"})
		return
	}

	utils.WriteSuccessJSON(w, project)
}

// ListProjects godoc
// @Summary List all projects
// @Description Retrieves an array of all existing projects
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response{data=[]models.Project}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects [get]
func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Repo.List()
	if err != nil {
		log.Printf("Error retrieving projects: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve projects"})
		return
	}
	utils.WriteSuccessJSON(w, projects)
}

// GetProject godoc
// @Summary Get a project by ID
// @Description Retrieves a single project using the provided ID
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Project ID"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{id} [get]
func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
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

	utils.WriteSuccessJSON(w, project)
}

type updateData struct {
	ChannelID   string `json:"channel_id"`
	Description string `json:"description"`
}

// UpdateProject godoc
// @Summary Update a project by ID
// @Description Updates a project's details using the provided ID and request body
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Project ID"
// @Param project body updateData true "Updated project data"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{id} [put]
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
		return
	}

	var updateData updateData
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve project"})
		return
	}

	project.ChannelID = updateData.ChannelID
	project.Description = updateData.Description
	project.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := h.Repo.Update(project); err != nil {
		log.Printf("Error updating project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to update project"})
		return
	}

	utils.WriteSuccessJSON(w, project)
}

// DeleteProject godoc
// @Summary Delete a project by ID
// @Description Deletes a project using the provided ID
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Project ID"
// @Success 200 {object} utils.Response{data=map[string]string}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to delete project"})
		return
	}

	utils.WriteSuccessJSON(w, map[string]string{"message": "Project deleted successfully"})
}

// SendMessageToChannel godoc
// @Summary      Send a message to a Discord channel
// @Description  Sends a message to the Discord channel associated with the specified project ID. The message is rendered using a Go template provided in the request body.
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        id path string true "Project ID"
// @Param        message body models.DiscordMessage true "Message payload with template"
// @Success      200 {object} utils.Response{data=map[string]string}
// @Failure      400 {object} utils.Response{error=[]string}
// @Failure      500 {object} utils.Response{error=[]string}
// @Router       /api/v1/projects/{id}/send-message [post]
func (h *ProjectHandler) SendMessageToChannel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
		return
	}

	var message models.DiscordMessage
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting project: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve project"})
		return
	}

	discordService, err := services.NewDiscordService()
	if err != nil {
		log.Printf("Error initializing Discord service: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to initialize Discord service"})
		return
	}

	tmpl, err := template.New("message").Parse(message.Template)
	if err != nil {
		log.Printf("Error parsing message template: %v", err)
		utils.WriteBadRequestJSON(w, []string{"Invalid template formate"})
		return
	}

	var renderedMessage strings.Builder
	if err := tmpl.Execute(&renderedMessage, project); err != nil {
		log.Printf("Error rendering message template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to render message template"})
		return
	}

	if err := discordService.SendMessage(project.ChannelID, renderedMessage.String()); err != nil {
		log.Printf("Error sending message: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to send message"})
		return
	}

	utils.WriteSuccessJSON(w, map[string]string{"message": "Message sent successfully"})
}
