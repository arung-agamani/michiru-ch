package models

type Project struct {
	ID            string `json:"id" db:"id"`
	ProjectName   string `json:"project_name" db:"project_name"`
	ChannelID     string `json:"channel_id" db:"channel_id"`
	AddedBy       string `json:"added_by" db:"added_by"`
	CreatedAt     string `json:"created_at" db:"created_at"`
	UpdatedAt     string `json:"updated_at" db:"updated_at"`
	Description   string `json:"description,omitempty" db:"description"`
	WebhookOrigin string `json:"webhook_origin,omitempty" db:"webhook_origin"`
	WebhookURL    string `json:"webhook_url,omitempty" db:"webhook_url"`
	WebhookSecret string `json:"webhook_secret,omitempty" db:"webhook_secret"`
}
