package config

import (
	"path/filepath"
	"testing"
)

func TestReadWriteConfig(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "testconfig.json")

	// Override config file path for testing
	originalPathFunc := getConfigFilePath
	getConfigFilePath = func() (string, error) {
		return tmpFile, nil
	}
	defer func() { getConfigFilePath = originalPathFunc }()

	expected := Config{
		DBURL:           "postgres://localhost:5432/db",
		CurrentUserName: "testuser",
	}

	// Write config
	err := write(expected)
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	// Read config
	actual, err := Read()
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestSetUser(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "testconfig.json")

	// Override config file path for testing
	originalPathFunc := getConfigFilePath
	getConfigFilePath = func() (string, error) {
		return tmpFile, nil
	}
	defer func() { getConfigFilePath = originalPathFunc }()

	// Create initial config
	initial := Config{
		DBURL:           "postgres://localhost:5432/db",
		CurrentUserName: "olduser",
	}
	err := write(initial)
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	cfg, err := Read()
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	newUser := "newuser"
	err = cfg.SetUser(newUser)
	if err != nil {
		t.Fatalf("failed to set user: %v", err)
	}

	// Re-read to confirm it's saved
	updated, err := Read()
	if err != nil {
		t.Fatalf("failed to re-read config: %v", err)
	}

	if updated.CurrentUserName != newUser {
		t.Errorf("expected CurrentUserName to be %s, got %s", newUser, updated.CurrentUserName)
	}
}
