package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// Parses command-line arguments correctly when valid flags are provided
func TestGetConfigWithValidFlags(t *testing.T) {
	os.Args = []string{"cmd", "-filename", "test.txt", "-timeout", "5s"}
	expectedConfig := Config{Filename: "test.txt", Timeout: 5 * time.Second, Help: false}

	config, err := GetConfig()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("Expected config %v, got %v", expectedConfig, config)
	}
}

// Handles invalid flag gracefully by returning an error
func TestGetConfigWithInvalidFlag(t *testing.T) {
	os.Args = []string{"cmd", "-invalidflag"}

	_, err := GetConfig()

	if err == nil {
		t.Fatal("Expected an error for invalid flag, got nil")
	}

	expectedError := "invalid flag"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}
}

// Handles invalid flag values gracefully with an error
func TestGetConfigHandlesInvalidFlagValues(t *testing.T) {
	os.Args = []string{"cmd", "-filename"}
	_, err := GetConfig()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	expectedError := "invalid flag value: -filename"
	if err.Error() != expectedError {
		t.Errorf("expected error %v, got %v", expectedError, err.Error())
	}
}
