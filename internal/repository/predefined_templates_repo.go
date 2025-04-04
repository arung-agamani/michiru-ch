package repository

import (
	"michiru/internal/models"

	"github.com/jmoiron/sqlx"
)

type PredefinedTemplateRepository struct {
	DB *sqlx.DB
}

func NewPredefinedTemplateRepository(db *sqlx.DB) *PredefinedTemplateRepository {
	return &PredefinedTemplateRepository{DB: db}
}

func (r *PredefinedTemplateRepository) GetAll() ([]models.PredefinedTemplate, error) {
	query := `SELECT id, event_type, template, description, created_at, updated_at FROM predefined_templates`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []models.PredefinedTemplate
	for rows.Next() {
		var template models.PredefinedTemplate
		if err := rows.Scan(&template.ID, &template.EventType, &template.Template, &template.Description, &template.CreatedAt, &template.UpdatedAt); err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}
	return templates, nil
}

func (r *PredefinedTemplateRepository) Add(template *models.PredefinedTemplate) error {
	query := `INSERT INTO predefined_templates (event_type, template, description) 
              VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, template.EventType, template.Template, template.Description).
		Scan(&template.ID)
}

func (r *PredefinedTemplateRepository) Update(template *models.PredefinedTemplate) error {
	query := `UPDATE predefined_templates 
              SET event_type = $1, template = $2, description = $3, updated_at = CURRENT_TIMESTAMP 
              WHERE id = $4`
	_, err := r.DB.Exec(query, template.EventType, template.Template, template.Description, template.ID)
	return err
}

func (r *PredefinedTemplateRepository) Delete(id int) error {
	query := `DELETE FROM predefined_templates WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
