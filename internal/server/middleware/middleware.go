package middleware

import (
	"context"
	"michiru/internal/server/handlers/auth"
	"net/http"

	"github.com/coreos/go-oidc"
)

var verifier *oidc.IDTokenVerifier

func Init(v *oidc.IDTokenVerifier) {
	verifier = v
}

type contextKey string

const idTokenKey contextKey = "id_token"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Failed to read session cookie: "+err.Error(), http.StatusUnauthorized)
			return
		}

		idToken, exists := auth.GetSession(cookie.Value)
		if !exists {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.Background()
		token, err := verifier.Verify(ctx, idToken)
		if err != nil {
			http.Error(w, "Failed to verify ID token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(r.Context(), idTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
