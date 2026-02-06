package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	ServiceAccountPath string
	ProjectID          string
	SubscriptionID     string
}

// LoadConfig parses command line arguments and environment variables
func LoadConfig() (*Config, error) {
	var serviceAccountPath string
	var subscriptionID string

	// Check environment variable first
	envPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	// Parse command line flags
	flag.StringVar(&serviceAccountPath, "key", envPath, "Path to Google Cloud Service Account JSON file")
	flag.StringVar(&subscriptionID, "subscription", "", "Google Cloud Pub/Sub Subscription ID")
	flag.Parse()

	if serviceAccountPath == "" {
		return nil, errors.New("service account path is required. Use --key flag or GOOGLE_APPLICATION_CREDENTIALS env var")
	}

	// Validate file exists and is valid JSON
	projectID, err := validateServiceAccountFile(serviceAccountPath)
	if err != nil {
		return nil, fmt.Errorf("invalid service account file: %w", err)
	}

	return &Config{
		ServiceAccountPath: serviceAccountPath,
		ProjectID:          projectID,
		SubscriptionID:     subscriptionID,
	}, nil
}

func validateServiceAccountFile(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", errors.New("file does not exist")
	}
	if fileInfo.IsDir() {
		return "", errors.New("path is a directory, not a file")
	}

	// Read file and check if it's valid JSON
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if !json.Valid(content) {
		return "", errors.New("file content is not valid JSON")
	}

	// Try to extract project_id
	var creds struct {
		ProjectID string `json:"project_id"`
	}
	if err := json.Unmarshal(content, &creds); err != nil {
		return "", nil // Valid JSON but maybe not SA key, return empty project ID
	}

	return creds.ProjectID, nil
}
