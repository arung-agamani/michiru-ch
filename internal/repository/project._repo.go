package repository

import (
	"github.com/jmoiron/sqlx"

	"michiru/internal/models"
)

type ProjectRepository struct {
	DB *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) Insert(project *models.Project) error {
	_, err := r.DB.Exec(
		"INSERT INTO projects (id, project_name, channel_id, added_by, created_at, updated_at, description) VALUES ($1, $2, $3, $4, NOW(), NOW(), $5)",
		project.ID, project.ProjectName, project.ChannelID, project.AddedBy, project.Description,
	)
	return err
}

func (r *ProjectRepository) GetByID(id string) (*models.Project, error) {
	var project models.Project
	err := r.DB.Get(&project, "SELECT id, project_name, channel_id, added_by, created_at, updated_at, description, webhook_url, webhook_origin FROM projects WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) Update(project *models.Project) error {
	_, err := r.DB.Exec(
		"UPDATE projects SET channel_id=$1, description=$2, updated_at=NOW() WHERE id=$3",
		project.ChannelID, project.Description, project.ID,
	)
	return err
}

func (r *ProjectRepository) UpdateWebhook(project *models.Project) error {
	_, err := r.DB.Exec(
		"UPDATE projects SET webhook_origin=$1, webhook_url=$2, webhook_secret=$3 WHERE id=$4",
		project.WebhookOrigin, project.WebhookURL, project.WebhookSecret, project.ID,
	)
	return err
}

func (r *ProjectRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM projects WHERE id=$1", id)
	return err
}

func (r *ProjectRepository) List() ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.Select(&projects, "SELECT id, project_name, channel_id, added_by, created_at, updated_at, description, webhook_origin, webhook_url FROM projects")
	if err != nil {
		return nil, err
	}
	return projects, nil
}
