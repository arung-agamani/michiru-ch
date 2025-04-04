package handlers

import (
	"encoding/json"
	"michiru/internal/models"
	"michiru/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type TemplateHandler struct {
	TemplateRepo *repository.TemplateRepository
}

func NewTemplateHandler(templateRepo *repository.TemplateRepository) *TemplateHandler {
	return &TemplateHandler{TemplateRepo: templateRepo}
}

// GetTemplates godoc
// @Summary Get all templates for a project
// @Description Retrieve all templates associated with a specific project
// @Tags Templates
// @Param projectID path string true "Project ID"
// @Success 200 {array} models.Template
// @Failure 500 {object} map[string]string
// @Router /api/v1/projects/{projectID}/templates [get]
func (h *TemplateHandler) GetTemplates(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]

	templates, err := h.TemplateRepo.GetByProjectID(projectID)
	if err != nil {
		http.Error(w, "Failed to fetch templates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templates)
}

// AddTemplate godoc
// @Summary Add a new template
// @Description Create a new template for a project
// @Tags Templates
// @Param template body models.Template true "Template data"
// @Success 201 {object} models.Template
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/projects/{projectID}/templates [post]
func (h *TemplateHandler) AddTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.TemplateRepo.Create(&template); err != nil {
		http.Error(w, "Failed to create template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(template)
}

// UpdateTemplate godoc
// @Summary Update an existing template
// @Description Update the details of an existing template
// @Tags Templates
// @Param templateID path string true "Template ID"
// @Param template body models.Template true "Updated template data"
// @Success 200 {object} models.Template
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/templates/{templateID} [put]
func (h *TemplateHandler) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.TemplateRepo.Update(&template); err != nil {
		http.Error(w, "Failed to update template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(template)
}

// DeleteTemplate godoc
// @Summary Delete a template
// @Description Delete a template by its ID
// @Tags Templates
// @Param templateID path string true "Template ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/v1/templates/{templateID} [delete]
func (h *TemplateHandler) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	templateID := mux.Vars(r)["templateID"]

	if err := h.TemplateRepo.Delete(templateID); err != nil {
		http.Error(w, "Failed to delete template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
