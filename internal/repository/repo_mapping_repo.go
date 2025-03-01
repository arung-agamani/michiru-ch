package repository

import (
	"michiru/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoMappingRepository struct {
	DB *sqlx.DB
}

func NewRepoMappingRepository(db *sqlx.DB) *RepoMappingRepository {
	return &RepoMappingRepository{DB: db}
}

func (r *RepoMappingRepository) Insert(mapping *models.RepoMapping) error {
	_, err := r.DB.Exec(
		"INSERT INTO repo_mappings (id, repo_name, channel_id, added_by, created_at, updated_at, description) VALUES ($1, $2, $3, $4, NOW(), NOW(), $5)",
		mapping.ID, mapping.RepoName, mapping.ChannelID, mapping.AddedBy, mapping.Description,
	)
	return err
}

func (r *RepoMappingRepository) GetByID(id string) (*models.RepoMapping, error) {
	var mapping models.RepoMapping
	err := r.DB.QueryRow("SELECT id, repo_name, channel_id, added_by, created_at, updated_at, description FROM repo_mappings WHERE id=$1", id).
		Scan(&mapping.ID, &mapping.RepoName, &mapping.ChannelID, &mapping.AddedBy, &mapping.CreatedAt, &mapping.UpdatedAt, &mapping.Description)

	if err != nil {
		return nil, err
	}
	return &mapping, nil
}