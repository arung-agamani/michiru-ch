package repository

import (
	"michiru/internal/models"

	"github.com/jmoiron/sqlx"
)

type TemplateRepository struct {
	DB *sqlx.DB
}

func NewTemplateRepository(db *sqlx.DB) *TemplateRepository {
	return &TemplateRepository{DB: db}
}

func (r *TemplateRepository) Insert(template *models.Template) error {
	query := `INSERT INTO templates (project_id, event_type, template, description) 
              VALUES ($1, $2, $3, $4) 
              ON CONFLICT (project_id, event_type) 
              DO UPDATE SET template = EXCLUDED.template, description = EXCLUDED.description, updated_at = CURRENT_TIMESTAMP 
              RETURNING id`
	return r.DB.QueryRow(query, template.ProjectID, template.EventType, template.Template, template.Description).
		Scan(&template.ID)
}

func (r *TemplateRepository) GetByProjectID(projectID string) ([]models.Template, error) {
	query := `SELECT id, project_id, event_type, template, description, created_at, updated_at 
              FROM templates WHERE project_id = $1`
	rows, err := r.DB.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []models.Template
	for rows.Next() {
		var template models.Template
		if err := rows.Scan(&template.ID, &template.ProjectID, &template.EventType, &template.Template, &template.Description, &template.CreatedAt, &template.UpdatedAt); err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}
	return templates, nil
}

func (r *TemplateRepository) GetByID(templateID string) (*models.Template, error) {
	query := `SELECT id, project_id, event_type, template, description, created_at, updated_at 
			  FROM templates WHERE id = $1`
	var template models.Template
	err := r.DB.Get(&template, query, templateID)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *TemplateRepository) Update(template *models.Template) error {
	query := `UPDATE templates 
              SET event_type = $1, template = $2, description = $3, updated_at = CURRENT_TIMESTAMP 
              WHERE id = $4`
	_, err := r.DB.Exec(query, template.EventType, template.Template, template.Description, template.ID)
	return err
}

func (r *TemplateRepository) Delete(templateID string) error {
	query := `DELETE FROM templates WHERE id = $1`
	_, err := r.DB.Exec(query, templateID)
	return err
}
