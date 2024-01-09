// utils/authenticator/firebase.go

package authenticator

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

// FirebaseManager handles Firebase authentication and database operations.
type FirebaseManager struct {
	authClient *auth.Client
	dbClient   *db.Client
}

// NewFirebaseManager creates a new instance of FirebaseManager.
func NewFirebaseManager(credentialsFilePath string, databaseURL string) (*FirebaseManager, error) {
	opt := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase Auth client: %v", err)
	}

	dbClient, err := app.DatabaseWithURL(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase Database client: %v", err)
	}

	return &FirebaseManager{
		authClient: authClient,
		dbClient:   dbClient,
	}, nil
}

// AuthenticateUser authenticates a user using Firebase Auth.
func (fm *FirebaseManager) AuthenticateUser(email, password string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	user, err := fm.authClient.CreateUser(context.Background(), params)
	if err != nil {
		return "", fmt.Errorf("failed to authenticate user: %v", err)
	}

	token, err := fm.authClient.CustomToken(context.Background(), user.UID)
	if err != nil {
		return "", fmt.Errorf("failed to generate custom token: %v", err)
	}

	return token, nil
}

// AuthorizeUserAccess authorizes access based on the user's UID.
func (fm *FirebaseManager) AuthorizeUserAccess(uid string) error {
	// Your authorization logic here
	// Example: Check user permissions or roles in Firebase Realtime Database
	return nil
}

// SendFlowData sends flow data to Firebase.
func (fm *FirebaseManager) SendFlowData(data map[string]interface{}) error {
	// Your logic for sending flow data to Firebase
	return nil
}
