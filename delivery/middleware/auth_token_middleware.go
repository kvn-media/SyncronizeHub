// delivery/middleware/auth_token_middleware.go

package middleware

import (
	"net/http"

	"github.com/kvn-media/SyncronizeHub/utils/authenticator"
)

// AuthTokenMiddleware is a middleware for authentication token validation.
type AuthTokenMiddleware struct {
	AccessToken authenticator.AccessToken
}

// NewAuthTokenMiddleware creates a new instance of AuthTokenMiddleware.
func NewAuthTokenMiddleware(accessToken authenticator.AccessToken) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		AccessToken: accessToken,
	}
}

// Middleware function for validating authentication token.
func (m *AuthTokenMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromRequest(r)

		if tokenString == "" {
			m.respondWithError(w, http.StatusUnauthorized, "Authorization token is missing")
			return
		}

		claims, err := m.AccessToken.VerifyToken(tokenString)
		if err != nil {
			m.respondWithError(w, http.StatusUnauthorized, "Invalid authorization token")
			return
		}

		// Attach user claims to the request context for later use
		r = attachClaimsToContext(r, claims)

		// Continue to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

// Utility function to extract token from the request (placeholder implementation).
func extractTokenFromRequest(r *http.Request) string {
	// Logic to extract token, e.g., from headers or cookies
	return ""
}

// Utility function to attach user claims to the request context.
func attachClaimsToContext(r *http.Request, claims authenticator.UserClaims) *http.Request {
	// Attach claims to the request context
	// Example: context.WithValue(r.Context(), "userClaims", claims)
	return r
}

// Utility function to send a standardized JSON error response.
func (m *AuthTokenMiddleware) respondWithError(w http.ResponseWriter, code int, message string) {
	response := map[string]interface{}{"error": message}
	respondWithJSON(w, code, response)
}

// Utility function to send a standardized JSON response.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Implementation omitted for brevity (similar to previous example)
}
