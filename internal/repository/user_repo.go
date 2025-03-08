package repository

import (
	"michiru/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Insert(user *models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (id, username, email, created_at, api_token) VALUES ($1, $2, $3, NOW(), $4)",
		user.ID, user.Username, user.Email, user.APIToken,
	)
	return err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Get(&user, "SELECT id, username, email, created_at, api_token FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SetAPIToken(email, token string) (*models.User, error) {
	_, err := r.DB.Exec("UPDATE users SET api_token=$1 WHERE email=$2", token, email)
	if err != nil {
		return nil, err
	}

	return r.GetByEmail(email)
}
