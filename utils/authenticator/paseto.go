// utils/authenticator/paseto.go

package authenticator

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

// PasetoManager handles PASETO token operations.
type PasetoManager struct {
	symmetricKey []byte
}

// NewPasetoManager creates a new instance of PasetoManager.
func NewPasetoManager(symmetricKey []byte) *PasetoManager {
	return &PasetoManager{
		symmetricKey: symmetricKey,
	}
}

// GenerateToken generates a PASETO token for the given subject and expiration time.
func (pm *PasetoManager) GenerateToken(subject string, expiration time.Time) (string, error) {
	claims := map[string]interface{}{
		"sub": subject,
		"exp": expiration.Unix(),
	}

	token, err := paseto.NewV2().Encrypt(pm.symmetricKey, claims, nil)
	if err != nil {
		return "", fmt.Errorf("error generating PASETO token: %v", err)
	}

	return token, nil
}

// VerifyToken verifies the authenticity of a PASETO token.
func (pm *PasetoManager) VerifyToken(token string) (string, error) {
	var claims map[string]interface{}
	err := paseto.NewV2().Decrypt(token, pm.symmetricKey, &claims, nil)
	if err != nil {
		return "", fmt.Errorf("error verifying PASETO token: %v", err)
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("missing or invalid subject in PASETO token")
	}

	return subject, nil
}
