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
	var templates []models.PredefinedTemplate
	err := r.DB.Select(&templates, `SELECT id, event_type, template, description, created_at, updated_at FROM predefined_templates`)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (r *PredefinedTemplateRepository) GetByID(id int) (*models.PredefinedTemplate, error) {
	var template models.PredefinedTemplate
	err := r.DB.Get(&template, `SELECT id, event_type, template, description, created_at, updated_at FROM predefined_templates WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *PredefinedTemplateRepository) Insert(template *models.PredefinedTemplate) error {
	return r.DB.QueryRow(
		`INSERT INTO predefined_templates (event_type, template, description, created_at, updated_at) 
         VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`,
		template.EventType, template.Template, template.Description,
	).Scan(&template.ID)
}

func (r *PredefinedTemplateRepository) Update(template *models.PredefinedTemplate) error {
	_, err := r.DB.Exec(
		`UPDATE predefined_templates 
         SET event_type = $1, template = $2, description = $3, updated_at = NOW() 
         WHERE id = $4`,
		template.EventType, template.Template, template.Description, template.ID,
	)
	return err
}

func (r *PredefinedTemplateRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM predefined_templates WHERE id = $1`, id)
	return err
}
