package handlers

import (
	"encoding/json"
	"log"
	"michiru/internal/models"
	"michiru/internal/repository"
	"michiru/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TemplateHandler struct {
	Repo repository.TemplateRepository
}

func NewTemplateHandler(repo repository.TemplateRepository) *TemplateHandler {
	return &TemplateHandler{Repo: repo}
}

// GetTemplates godoc
// @Summary Get all templates for a project
// @Description Retrieve all templates associated with a specific project
// @Tags Templates
// @Param projectID path string true "Project ID"
// @Success 200 {object} utils.Response{data=[]models.Template}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{projectID}/templates [get]
func (h *TemplateHandler) GetTemplates(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]

	templates, err := h.Repo.GetByProjectID(projectID)
	if err != nil {
		log.Printf("Error fetching templates: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to fetch templates"})
		return
	}

	if templates == nil {
		templates = []models.Template{}
	}

	utils.WriteSuccessJSON(w, templates)
}

// AddTemplate godoc
// @Summary Add a new template
// @Description Create a new template for a project
// @Tags Templates
// @Param template body models.Template true "Template data"
// @Success 201 {object} utils.Response{data=models.Template}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/projects/{projectID}/templates [post]
func (h *TemplateHandler) AddTemplate(w http.ResponseWriter, r *http.Request) {
	supportedEventTypes := []string{"push", "pull_request", "issue_comment", "release", "workflow_dispatch"}

	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	if !isValidEventType(template.EventType, supportedEventTypes) {
		utils.WriteBadRequestJSON(w, []string{"Unsupported event type. Supported types are: push, pull_request, issue_comment, release, workflow_dispatch"})
		return
	}

	if err := h.Repo.Insert(&template); err != nil {
		log.Printf("Error creating template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to create template"})
		return
	}

	utils.WriteSuccessJSON(w, template)
}

func isValidEventType(eventType string, supportedEventTypes []string) bool {
	for _, supportedType := range supportedEventTypes {
		if eventType == supportedType {
			return true
		}
	}
	return false
}

// UpdateTemplate godoc
// @Summary Update an existing template
// @Description Update the details of an existing template
// @Tags Templates
// @Param templateID path string true "Template ID"
// @Param template body models.Template true "Updated template data"
// @Success 200 {object} utils.Response{data=models.Template}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/templates/{templateID} [put]
func (h *TemplateHandler) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateID := vars["templateID"]

	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	id, err := strconv.Atoi(templateID)
	if err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid template ID"})
		return
	}
	template.ID = id
	if err := h.Repo.Update(&template); err != nil {
		log.Printf("Error updating template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to update template"})
		return
	}

	utils.WriteSuccessJSON(w, template)
}

// DeleteTemplate godoc
// @Summary Delete a template
// @Description Delete a template by its ID
// @Tags Templates
// @Param templateID path string true "Template ID"
// @Success 204
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/templates/{templateID} [delete]
func (h *TemplateHandler) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateID := vars["templateID"]

	if err := h.Repo.Delete(templateID); err != nil {
		log.Printf("Error deleting template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to delete template"})
		return
	}

	utils.WriteSuccessJSON(w, map[string]string{"message": "Template deleted successfully"})
}
