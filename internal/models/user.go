package models

type User struct {
	ID             string  `json:"id" db:"id"`
	Username       string  `json:"username" db:"username"`
	Email          string  `json:"email" db:"email"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	APIToken       *string `json:"api_token,omitempty" db:"api_token"`
	Name           *string `json:"name,omitempty" db:"name"`
	ProfilePicture *string `json:"profile_picture,omitempty" db:"profile_picture"`
}
