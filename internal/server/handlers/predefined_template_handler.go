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

type PredefinedTemplateHandler struct {
	Repo repository.PredefinedTemplateRepository
}

func NewPredefinedTemplateHandler(repo repository.PredefinedTemplateRepository) *PredefinedTemplateHandler {
	return &PredefinedTemplateHandler{Repo: repo}
}

// GetPredefinedTemplates godoc
// @Summary Get all predefined templates
// @Description Retrieve all predefined templates
// @Tags PredefinedTemplates
// @Success 200 {object} utils.Response{data=[]models.PredefinedTemplate}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/predefined-templates [get]
func (h *PredefinedTemplateHandler) GetPredefinedTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := h.Repo.GetAll()
	if err != nil {
		log.Printf("Error fetching predefined templates: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to fetch predefined templates"})
		return
	}

	if templates == nil {
		templates = []models.PredefinedTemplate{}
	}

	utils.WriteSuccessJSON(w, templates)
}

// GetPredefinedTemplateByID godoc
// @Summary Get a predefined template by ID
// @Description Retrieve a predefined template by its ID
// @Tags PredefinedTemplates
// @Param templateID path string true "Template ID"
// @Success 200 {object} utils.Response{data=models.PredefinedTemplate}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 404 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/predefined-templates/{templateID} [get]
func (h *PredefinedTemplateHandler) GetPredefinedTemplateByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateID := vars["templateID"]
	id, err := strconv.Atoi(templateID)
	if err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid template ID"})
		return
	}

	template, err := h.Repo.GetByID(id)
	if err != nil {
		log.Printf("Error fetching predefined template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to fetch predefined template"})
		return
	}

	if template == nil {
		utils.WriteNotFoundJSON(w, []string{"Predefined template not found"})
		return
	}

	utils.WriteSuccessJSON(w, template)
}

type AddPredefinedTemplateRequest struct {
	EventType   string `json:"event_type"`
	Template    string `json:"template"`
	Description string `json:"description"`
}

// AddPredefinedTemplate godoc
// @Summary Add a new predefined template
// @Description Create a new predefined template
// @Tags PredefinedTemplates
// @Param template body AddPredefinedTemplateRequest true "Predefined template data"
// @Success 201 {object} utils.Response{data=models.PredefinedTemplate}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/predefined-templates [post]
func (h *PredefinedTemplateHandler) AddPredefinedTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.PredefinedTemplate
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	if err := h.Repo.Insert(&template); err != nil {
		log.Printf("Error inserting predefined template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to create predefined template"})
		return
	}

	utils.WriteSuccessJSON(w, template)
}

// UpdatePredefinedTemplate godoc
// @Summary Update an existing predefined template
// @Description Update the details of an existing predefined template
// @Tags PredefinedTemplates
// @Param templateID path string true "Template ID"
// @Param template body models.PredefinedTemplate true "Updated predefined template data"
// @Success 200 {object} utils.Response{data=models.PredefinedTemplate}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/predefined-templates/{templateID} [put]
func (h *PredefinedTemplateHandler) UpdatePredefinedTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateID := vars["templateID"]
	id, err := strconv.Atoi(templateID)
	if err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid template ID"})
		return
	}

	var template models.PredefinedTemplate
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	template.ID = id
	if err := h.Repo.Update(&template); err != nil {
		log.Printf("Error updating predefined template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to update predefined template"})
		return
	}

	utils.WriteSuccessJSON(w, template)
}

// DeletePredefinedTemplate godoc
// @Summary Delete a predefined template
// @Description Delete a predefined template by its ID
// @Tags PredefinedTemplates
// @Param templateID path string true "Template ID"
// @Success 204
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /api/v1/predefined-templates/{templateID} [delete]
func (h *PredefinedTemplateHandler) DeletePredefinedTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateID := vars["templateID"]
	id, err := strconv.Atoi(templateID)
	if err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid template ID"})
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		log.Printf("Error deleting predefined template: %v", err)
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to delete predefined template"})
		return
	}

	utils.WriteSuccessJSON(w, map[string]string{"message": "Predefined template deleted successfully"})
}
