package reader

import (
	"os"
	"testing"
)

// Successfully reads a CSV file and returns its contents as a slice of strings
func TestReadCSVSuccess(t *testing.T) {
	filename := "test.csv"
	expected := []string{"q1,ans1", "q2,ans2", "q3,ans3"}

	err := os.WriteFile(filename, []byte("q1,ans1\nq2,ans2\nq3,ans3"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	result, err := ReadCSV(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(result))
	}

	for i, line := range expected {
		if result[i] != line {
			t.Errorf("Expected line %d to be %s, got %s", i, line, result[i])
		}
	}
}

// Handles non-existent file paths gracefully by returning an error

func TestReadCSVFileNotFound(t *testing.T) {
	filename := "non_existent.csv"

	_, err := ReadCSV(filename)
	if err == nil {
		t.Fatal("Expected an error for non-existent file, got nil")
	}
}
