// utils/authenticator/firebase.go

package authenticator

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FirebaseAuthenticator represents the Firebase authentication utility.
type FirebaseAuthenticator struct {
	app *firebase.App
}

// NewFirebaseAuthenticator initializes a new Firebase authenticator.
func NewFirebaseAuthenticator(ctx context.Context, credentialsFile string) (*FirebaseAuthenticator, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	return &FirebaseAuthenticator{app: app}, nil
}

// GeneratePasetoToken creates a PASETO token for the user.
func (fa *FirebaseAuthenticator) GeneratePasetoToken(uid string) (string, error) {
	// Fetch user information from Firebase
	user, err := fa.getUser(uid)
	if err != nil {
		return "", fmt.Errorf("error fetching user from Firebase: %v", err)
	}

	// Create PASETO token claims
	claims := map[string]interface{}{
		"uid":   uid,
		"email": user.Email,
		// Add other relevant claims as needed
	}

	// Sign the PASETO token using a private key (replace with your implementation)
	privateKey := getPrivateKey()

	pasetoAuthenticator := NewPasetoAuthenticator(privateKey, nil) // Use the appropriate public key

	// Generate PASETO token
	token, err := pasetoAuthenticator.GenerateToken(claims)
	if err != nil {
		return "", fmt.Errorf("error generating PASETO token: %v", err)
	}

	return token, nil
}

// VerifyPasetoToken verifies the PASETO token and returns the user claims.
func (fa *FirebaseAuthenticator) VerifyPasetoToken(token string) (map[string]interface{}, error) {
	// Verify the PASETO token using the public key (replace with your implementation)
	publicKey := getPublicKey()

	pasetoAuthenticator := NewPasetoAuthenticator(nil, publicKey) // Use the appropriate private key

	// Verify PASETO token
	claims, err := pasetoAuthenticator.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("error verifying PASETO token: %v", err)
	}

	return claims, nil
}

// getUser retrieves user information from Firebase using the provided UID.
func (fa *FirebaseAuthenticator) getUser(uid string) (*auth.UserRecord, error) {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		log.Fatalf("error fetching user data: %v\n", err)
	}

	return user, nil
}

// getPrivateKey retrieves the private key for signing PASETO tokens.
// Replace this function with your actual private key retrieval logic.
func getPrivateKey() ed25519.PrivateKey {
	// Replace with your implementation
}

// getPublicKey retrieves the public key for verifying PASETO tokens.
// Replace this function with your actual public key retrieval logic.
func getPublicKey() ed25519.PublicKey {
	// Replace with your implementation
}
