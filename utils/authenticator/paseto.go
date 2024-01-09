// utils/authenticator/paseto.go

package authenticator

import (
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
	"time"
)

// PasetoAuthenticator represents the PASETO authentication utility.
type PasetoAuthenticator struct {
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

// NewPasetoAuthenticator initializes a new PASETO authenticator.
func NewPasetoAuthenticator(privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) *PasetoAuthenticator {
	return &PasetoAuthenticator{privateKey: privateKey, publicKey: publicKey}
}

// GenerateToken creates a PASETO token with the provided claims.
func (pa *PasetoAuthenticator) GenerateToken(claims map[string]interface{}) (string, error) {
	v2 := paseto.NewV2()
	now := time.Now()

	// Set expiration to 1 hour from now
	expiration := now.Add(time.Hour)

	token, err := v2.Encrypt(pa.privateKey, claims, paseto.Expiration(expiration))
	if err != nil {
		return "", err
	}

	return token, nil
}

// VerifyToken verifies the PASETO token and returns the claims.
func (pa *PasetoAuthenticator) VerifyToken(token string) (map[string]interface{}, error) {
	v2 := paseto.NewV2()

	var claims map[string]interface{}
	err := v2.Decrypt(token, pa.publicKey, &claims, nil)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
