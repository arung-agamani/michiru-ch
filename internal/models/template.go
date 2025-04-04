package models

import "time"

type Template struct {
	ID          int       `json:"id" db:"id"`
	ProjectID   string    `json:"project_id" db:"project_id"`             // Foreign key to Project
	EventType   string    `json:"event_type" db:"event_type"`             // Event type (e.g., "push", "pull_request")
	Template    string    `json:"template" db:"template"`                 // Template content
	Description *string   `json:"description,omitempty" db:"description"` // Optional description
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type PredefinedTemplate struct {
	ID          int       `json:"id" db:"id"`
	EventType   string    `json:"event_type" db:"event_type"`             // Event type (e.g., "push", "pull_request")
	Template    string    `json:"template" db:"template"`                 // Template content
	Description *string   `json:"description,omitempty" db:"description"` // Optional description
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
