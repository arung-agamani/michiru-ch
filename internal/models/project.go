package models

type Project struct {
	ID          string `json:"id" db:"id"`
	ProjectName string `json:"project_name" db:"project_name"`
	ChannelID   string `json:"channel_id" db:"channel_id"`
	AddedBy     string `json:"added_by" db:"added_by"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	Description string `json:"description,omitempty" db:"description"`
}
