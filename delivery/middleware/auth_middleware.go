// delivery/middleware/auth.go

package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"context"
	"github.com/kvn-media/SyncronizeHub/utils/authenticator"
)

// AuthenticationMiddleware is a middleware function to ensure authentication
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		tokenString := ExtractTokenFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		// Verify the token using the authenticator package
		userID, err := authenticator.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unauthorized: %v", err), http.StatusUnauthorized)
			return
		}

		// Attach the user ID to the request context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)

		// Proceed to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

// extractTokenFromHeader extracts the token from the Authorization header
func ExtractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Check if the Authorization header is in the format "Bearer <token>"
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ""
	}

	return headerParts[1]
}