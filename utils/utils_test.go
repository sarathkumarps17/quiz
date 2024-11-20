package utils

import (
	"testing"
	"time"
)

// Correctly parses valid time duration strings with units
func TestParseTimeDurationWithUnits(t *testing.T) {
	durationStr := "2h45m"
	expectedDuration, _ := time.ParseDuration("2h45m")

	result, err := ParseTimeDuration(durationStr)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != expectedDuration {
		t.Errorf("Expected %v, got %v", expectedDuration, result)
	}
}

// Handles empty string input gracefully
func TestParseTimeDurationEmptyString(t *testing.T) {
	durationStr := ""
	expectedDuration := time.Duration(0)

	result, err := ParseTimeDuration(durationStr)

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	if result != expectedDuration {
		t.Errorf("Expected %v, got %v", expectedDuration, result)
	}
}
