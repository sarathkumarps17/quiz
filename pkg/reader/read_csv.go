package reader

import (
	"encoding/csv"
	"os"
	"strings"
)

// ReadCSV reads the file at the given filename and returns a slice of strings
// representing its contents, split by newline. If the file cannot be read,
// an error is returned.
func ReadCSV(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Convert records to strings with line separator
	var lines []string
	for _, record := range records {
		lines = append(lines, strings.Join(record, ","))
	}

	return lines, nil
}
