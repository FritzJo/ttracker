package modules

import (
	m "github.com/FritzJo/ttracker/modules"
	"os"
	"testing"
)

func TestValidateCSVFile(t *testing.T) {
	// Create a temporary CSV file for testing
	file, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	// Write test data to the CSV file
	csvData := `type;date;start;end;overtime
R;2023-05-01;7:45;17:00;15
R;2023-05-02;7:40;17:09;29
R;2023-05-03;8:23;17:45;22
R;2023-05-04;7:50;16:45;-5
O;0001-01-01;;;-480`
	_, err = file.WriteString(csvData)
	if err != nil {
		t.Fatalf("Failed to write to CSV file: %v", err)
	}

	// Call the function being tested
	err = m.ValidateCSVFile(file.Name())

	// Assert the result
	if err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}
