package auth

import (
	"context"
	"encoding/base64"
	"log"
	"michiru/internal/models"
	"michiru/internal/repository"
	"michiru/internal/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	redirectURL  string
)

var (
	Provider     *oidc.Provider
	oauth2Config *oauth2.Config
	Verifier     *oidc.IDTokenVerifier
	userRepo     *repository.UserRepository
)

func Init(db *sqlx.DB) {
	clientID = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL = os.Getenv("REDIRECT_URL")
	ctx := context.Background()

	p, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     p.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	Provider = p
	Verifier = p.Verifier(&oidc.Config{ClientID: clientID})
	userRepo = repository.NewUserRepository(db)
	InitSessionStore(db)
}

func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, oauth2Config.AuthCodeURL("state"), http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to read session cookie: "+err.Error(), http.StatusUnauthorized)
		return
	}

	DeleteSession(cookie.Value)

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No ID token in response", http.StatusInternalServerError)
		return
	}

	idToken, err := Verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	email := claims["email"].(string)
	username := strings.Split(email, "@")[0]

	_, err = userRepo.GetByEmail(email)
	if err != nil {
		user := &models.User{
			ID:        uuid.New().String(),
			Username:  username,
			Email:     email,
			CreatedAt: time.Now().Format(time.RFC3339),
		}
		if err := userRepo.Insert(user); err != nil {
			http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	sessionToken := uuid.New().String()
	SetSession(sessionToken, rawIDToken)

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to read session cookie: "+err.Error(), http.StatusUnauthorized)
		return
	}

	idToken, exists := GetSession(cookie.Value)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	token, err := Verifier.Verify(ctx, idToken)
	if err != nil {
		http.Error(w, "Failed to verify ID token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	var claims map[string]any
	if err := token.Claims(&claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	email := claims["email"].(string)
	user, err := userRepo.GetByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSON(w, user)
}

// make handler that will generate an API token for the user
func GenerateAPIToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to read session cookie: "+err.Error(), http.StatusUnauthorized)
		return
	}

	idToken, exists := GetSession(cookie.Value)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	token, err := Verifier.Verify(ctx, idToken)
	if err != nil {
		http.Error(w, "Failed to verify ID token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	var claims map[string]any
	if err := token.Claims(&claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	email := claims["email"].(string)
	_, err = userRepo.GetByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	// generate a new API token in base64
	newAPIToken := base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))
	user, err := userRepo.SetAPIToken(email, newAPIToken)
	if err != nil {
		http.Error(w, "Failed to generate API token", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSON(w, user)
}
