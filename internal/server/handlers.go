package server

import (
	"encoding/json"
	"fmt"
	_ "michiru/docs"
	"michiru/internal/db"
	"michiru/internal/repository"
	"michiru/internal/server/handlers"
	"michiru/internal/server/handlers/auth"
	"michiru/internal/server/middleware"
	"michiru/internal/services"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Michiru API
// @version 1.0
// @description API documentation for Michiru
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @security BearerAuth

func RegisterRoutes(router *mux.Router) {
	dbConn, err := db.Connect()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	auth.Init(dbConn)
	middleware.Init(auth.Verifier, dbConn)

	projectRepo := repository.NewProjectRepository(dbConn)
	templateRepo := repository.NewTemplateRepository(dbConn)
	predefinedTemplateRepo := repository.NewPredefinedTemplateRepository(dbConn)

	projectHandler := handlers.NewProjectHandler(*projectRepo)
	projectWebhookHandler := handlers.NewProjectWebhookHandler(*projectRepo)
	templateHandler := handlers.NewTemplateHandler(*templateRepo)
	predefinedTemplateHandler := handlers.NewPredefinedTemplateHandler(*predefinedTemplateRepo)

	router.HandleFunc("/auth/login", auth.Login)
	router.HandleFunc("/auth/logout", auth.Logout)
	router.HandleFunc("/auth/callback", auth.Callback)
	router.HandleFunc("/auth/refresh", auth.RefreshToken)

	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.AuthMiddleware)

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(middleware.AuthMiddleware)
	authRouter.HandleFunc("/me", auth.Me).Methods("GET")
	authRouter.HandleFunc("/me/genereate-api-key", auth.GenerateAPIToken).Methods("POST")
	apiV1.HandleFunc("/discord", SendMessageHandler).Methods("POST")
	apiV1.HandleFunc("/github-webhook", handlers.HandleGithubWebhook)
	apiV1.HandleFunc("/projects", projectHandler.CreateProject).Methods("POST")
	apiV1.HandleFunc("/projects", projectHandler.ListProjects).Methods("GET")
	apiV1.HandleFunc("/projects/{id}", projectHandler.UpdateProject).Methods("PUT")
	apiV1.HandleFunc("/projects/{id}", projectHandler.DeleteProject).Methods("DELETE")
	apiV1.HandleFunc("/projects/{id}", projectHandler.GetProject).Methods("GET")
	apiV1.HandleFunc("/projects/{id}/send-message", projectHandler.SendMessageToChannel).Methods("POST")
	apiV1.HandleFunc("/projects/{id}/webhook", projectWebhookHandler.UpdateWebhook).Methods("PUT")
	apiV1.HandleFunc("/projects/{id}/webhook", projectWebhookHandler.GenerateWebhook).Methods("POST")
	apiV1.HandleFunc("/projects/{projectID}/templates", templateHandler.GetTemplates).Methods("GET")
	apiV1.HandleFunc("/projects/{projectID}/templates", templateHandler.AddTemplate).Methods("POST")

	apiV1.HandleFunc("/templates/{templateID}", templateHandler.UpdateTemplate).Methods("PUT")
	apiV1.HandleFunc("/templates/{templateID}", templateHandler.DeleteTemplate).Methods("DELETE")

	apiV1.HandleFunc("/predefined-templates", predefinedTemplateHandler.GetPredefinedTemplates).Methods("GET")
	apiV1.HandleFunc("/predefined-templates", predefinedTemplateHandler.AddPredefinedTemplate).Methods("POST")
	apiV1.HandleFunc("/predefined-templates/{templateID}", predefinedTemplateHandler.GetPredefinedTemplateByID).Methods("GET")
	apiV1.HandleFunc("/predefined-templates/{templateID}", predefinedTemplateHandler.UpdatePredefinedTemplate).Methods("PUT")
	apiV1.HandleFunc("/predefined-templates/{templateID}", predefinedTemplateHandler.DeletePredefinedTemplate).Methods("DELETE")

	// Publicly accessible endpoint for GitHub webhooks
	router.HandleFunc("/api/v1/projects/{id}/webhook/{webhookId}", projectWebhookHandler.HandleWebhookPayload).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/", HomeHandler)

	corsAllowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	var allowedOrigins []string
	if corsAllowedOrigins != "" {
		allowedOrigins = strings.Split(corsAllowedOrigins, ",")
	} else {
		allowedOrigins = []string{"localhost:5173", "localhost:8080", "https://michiru.howlingmoon.dev"}
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(corsHandler.Handler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Konnichiwa, sekai")
}

type MessageRequest struct {
	ChannelID string `json:"channel"`
	Message   string `json:"message"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	discordService, err := services.NewDiscordService()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error initializing Discord service: %v", err), http.StatusInternalServerError)
		return
	}
	defer discordService.Close()

	err = discordService.SendMessage(req.ChannelID, req.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending message: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}
