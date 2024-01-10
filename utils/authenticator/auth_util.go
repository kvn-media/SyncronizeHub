// utils/authenticator/auth_util.go

package authenticator

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTSecretKey is a secret key for signing and validating JWTs
var JWTSecretKey = []byte("your-secret-key") // Replace with your actual secret key

// GenerateToken generates a new JWT
func GenerateToken(userID string) (string, error) {
	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	// Sign the token with the secret key
	signedToken, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %v", err)
	}

	return signedToken, nil
}

// VerifyToken verifies the authenticity of a token
func VerifyToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("invalid or expired token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("invalid or expired token")
	}

	// Retrieve the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid user ID in token claims")
	}

	return userID, nil
}

// CleanupToken is a no-op in this version (no Redis cleanup)
func CleanupToken(token string) error {
	// No Redis cleanup needed in this version
	return nil
}
