package handlers

import (
	"encoding/json"
	"michiru/internal/models"
	"michiru/internal/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PredefinedTemplateHandler struct {
	PredefinedTemplateRepo *repository.PredefinedTemplateRepository
}

func NewPredefinedTemplateHandler(predefinedTemplateRepo *repository.PredefinedTemplateRepository) *PredefinedTemplateHandler {
	return &PredefinedTemplateHandler{PredefinedTemplateRepo: predefinedTemplateRepo}
}

// GetPredefinedTemplates godoc
// @Summary Get all predefined templates
// @Description Retrieve all predefined templates
// @Tags PredefinedTemplates
// @Success 200 {array} models.PredefinedTemplate
// @Failure 500 {object} map[string]string
// @Router /api/v1/predefined-templates [get]
func (h *PredefinedTemplateHandler) GetPredefinedTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := h.PredefinedTemplateRepo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch predefined templates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templates)
}

// AddPredefinedTemplate godoc
// @Summary Add a new predefined template
// @Description Create a new predefined template
// @Tags PredefinedTemplates
// @Param template body models.PredefinedTemplate true "Predefined template data"
// @Success 201 {object} models.PredefinedTemplate
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/predefined-templates [post]
func (h *PredefinedTemplateHandler) AddPredefinedTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.PredefinedTemplate
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.PredefinedTemplateRepo.Add(&template); err != nil {
		http.Error(w, "Failed to create predefined template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(template)
}

// UpdatePredefinedTemplate godoc
// @Summary Update an existing predefined template
// @Description Update the details of an existing predefined template
// @Tags PredefinedTemplates
// @Param templateID path string true "Template ID"
// @Param template body models.PredefinedTemplate true "Updated predefined template data"
// @Success 200 {object} models.PredefinedTemplate
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/predefined-templates/{templateID} [put]
func (h *PredefinedTemplateHandler) UpdatePredefinedTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.PredefinedTemplate
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.PredefinedTemplateRepo.Update(&template); err != nil {
		http.Error(w, "Failed to update predefined template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(template)
}

// DeletePredefinedTemplate godoc
// @Summary Delete a predefined template
// @Description Delete a predefined template by its ID
// @Tags PredefinedTemplates
// @Param templateID path string true "Template ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/v1/predefined-templates/{templateID} [delete]
func (h *PredefinedTemplateHandler) DeletePredefinedTemplate(w http.ResponseWriter, r *http.Request) {
	templateID := mux.Vars(r)["templateID"]
	id, err := strconv.Atoi(templateID)
	if err != nil {
		http.Error(w, "Invalid template ID", http.StatusBadRequest)
		return
	}

	if err := h.PredefinedTemplateRepo.Delete(id); err != nil {
		http.Error(w, "Failed to delete predefined template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
