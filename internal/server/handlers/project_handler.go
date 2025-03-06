package handlers

import (
	"encoding/json"
	"michiru/internal/models"
	"michiru/internal/repository"
	"michiru/internal/utils"
	"net/http"
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
// @Router /projects [post]
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
		return
	}

	if err := h.Repo.Insert(&project); err != nil {
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to create a project"})
		return
	}

	utils.WriteSuccessJSON(w, project)
}

// GetProject godoc
// @Summary Get a project by ID
// @Description Retrieves a single project using the provided ID
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id query string true "Project ID"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /projects [get]
func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
        return
    }

    project, err := h.Repo.GetByID(id)
    if err != nil {
        utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve project"})
        return
    }

    utils.WriteSuccessJSON(w, project)
}

// UpdateProject godoc
// @Summary Update a project by ID
// @Description Updates a project's details using the provided ID and request body
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id query string true "Project ID"
// @Param project body models.Project true "Updated project data"
// @Success 200 {object} utils.Response{data=models.Project}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /projects [put]
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        utils.WriteBadRequestJSON(w, []string{"Missing project ID"})
        return
    }

    var project models.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        utils.WriteBadRequestJSON(w, []string{"Invalid request payload"})
        return
    }

    project.ID = id
    if err := h.Repo.Update(&project); err != nil {
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
// @Param id query string true "Project ID"
// @Success 200 {object} utils.Response{data=map[string]string}
// @Failure 400 {object} utils.Response{error=[]string}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /projects [delete]
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
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

// ListProjects godoc
// @Summary List all projects
// @Description Retrieves an array of all existing projects
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response{data=[]models.Project}
// @Failure 500 {object} utils.Response{error=[]string}
// @Router /projects [get]
func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Repo.List()
	if err != nil {
		utils.WriteInternalServerErrorJSON(w, []string{"Failed to retrieve projects"})
		return
	}
	utils.WriteSuccessJSON(w, projects)
}