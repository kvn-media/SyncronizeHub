// delivery/middleware/auth_token_middleware.go

package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kvn-media/SyncronizeHub/utils/authenticator"
)

// AuthMiddleware is a middleware for authenticating requests using PASETO tokens.
type AuthMiddleware struct {
	pasetoManager *authenticator.PasetoManager
}

// NewAuthMiddleware creates a new instance of AuthMiddleware.
func NewAuthMiddleware(pasetoManager *authenticator.PasetoManager) *AuthMiddleware {
	return &AuthMiddleware{
		pasetoManager: pasetoManager,
	}
}

// Middleware function to authenticate requests using PASETO tokens.
func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request headers
		tokenString := extractTokenFromHeader(r)

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Verify the token
		subject, err := am.pasetoManager.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unauthorized: %v", err), http.StatusUnauthorized)
			return
		}

		// Attach the subject to the request context for further use
		ctx := WithUserID(r.Context(), subject)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// extractTokenFromHeader extracts the token from the Authorization header.
func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Expected header format: "Bearer <token>"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	return tokenParts[1]
}
