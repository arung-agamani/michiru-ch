package models

type RepoMapping struct {
	ID          string `json:"id"`
	RepoName    string `json:"repo_name"`
	ChannelID   string `json:"channel_id"`
	AddedBy     string `json:"added_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Description string `json:"description,omitempty"`
}