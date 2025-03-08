package middleware

import (
	"context"
	"michiru/internal/repository"
	"michiru/internal/server/handlers/auth"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/jmoiron/sqlx"
)

var verifier *oidc.IDTokenVerifier
var userRepo *repository.UserRepository

func Init(v *oidc.IDTokenVerifier, db *sqlx.DB) {
	verifier = v
	userRepo = repository.NewUserRepository(db)
	auth.Init(db)
}

type contextKey string

const idTokenKey contextKey = "id_token"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// http.Error(w, "Unauthorized", http.StatusUnauthorized)
				http.Redirect(w, r, "/auth/login", http.StatusFound)
				return
			}
			http.Error(w, "Failed to read session cookie: "+err.Error(), http.StatusUnauthorized)
			return
		}

		idToken, exists := auth.GetSession(cookie.Value)
		if !exists {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}

		ctx := context.Background()
		token, err := verifier.Verify(ctx, idToken)
		if err != nil {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}

		// Extract email from token claims
		var claims map[string]interface{}
		if err := token.Claims(&claims); err != nil {
			http.Error(w, "Failed to parse token claims: "+err.Error(), http.StatusUnauthorized)
			return
		}

		email := claims["email"].(string)
		user, err := userRepo.GetByEmail(email)
		if err != nil {
			http.Error(w, "User not found: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(r.Context(), idTokenKey, token)
		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
