package auth

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// TODO: figure out better way to store sessions
var db *sqlx.DB

func InitSessionStore(database *sqlx.DB) {
	db = database
}

func SetSession(key, value string) {
	_, err := db.Exec("INSERT INTO sessions (key, value) VALUES ($1, $2)", key, value)
	if err != nil {
		log.Printf("Failed to insert session: %v", err)
	}
}

func GetSession(key string) (string, bool) {
	var value string
	err := db.Get(&value, "SELECT value FROM sessions WHERE key=$1", key)
	if err != nil {
		return "", false
	}
	return value, true
}

func DeleteSession(key string) {
	_, err := db.Exec("DELETE FROM sessions WHERE key=$1", key)
	if err != nil {
		log.Printf("Failed to delete session: %v", err)
	}
}
