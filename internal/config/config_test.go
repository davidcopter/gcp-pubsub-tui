package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateServiceAccountFile(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Case 1: Valid JSON with project_id
	validFile := filepath.Join(tmpDir, "valid.json")
	err = os.WriteFile(validFile, []byte(`{"project_id": "my-project"}`), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Case 2: Invalid JSON
	invalidFile := filepath.Join(tmpDir, "invalid.json")
	err = os.WriteFile(invalidFile, []byte(`not json`), 0644)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		path     string
		wantErr  bool
		wantProj string
	}{
		{
			name:     "Valid File",
			path:     validFile,
			wantErr:  false,
			wantProj: "my-project",
		},
		{
			name:    "Invalid File",
			path:    invalidFile,
			wantErr: true,
		},
		{
			name:    "Non-existent File",
			path:    filepath.Join(tmpDir, "missing.json"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proj, err := validateServiceAccountFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateServiceAccountFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && proj != tt.wantProj {
				t.Errorf("validateServiceAccountFile() project = %v, want %v", proj, tt.wantProj)
			}
		})
	}
}
