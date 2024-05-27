package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/FritzJo/ttracker/modules/datatypes"
)

// ValidateCSVFile validates a CSV file at the given filepath against the expected format of the time tracker application.
// The CSV file should have the following columns: type, date (YYYY-MM-DD), start time (HH:MM), end time (HH:MM), and overtime (minutes).
// The function returns an error if the file is not valid, and nil if it is valid.
func ValidateCSVFile(filepath string) error {
	// Load the configuration file
	config, _ := datatypes.LoadConfig("config.json")

	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV reader and set properties
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 5
	csvReader.Comma = ';'
	csvReader.TrimLeadingSpace = true

	// Read all records in CSV file
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	// Loop through all records, starting from the second one
	for i, record := range records {
		if i != 0 {
			// Check if the record has the correct number of fields
			if len(record) != csvReader.FieldsPerRecord {
				return fmt.Errorf("line %d: invalid number of fields", i+1)
			}

			// Check if the date is in the correct format
			_, err := time.Parse("2006-01-02", record[1])
			if err != nil {
				return fmt.Errorf("line %d: invalid date format", i+1)
			}

			if record[0] == "O" {
				continue
			}

			// Check if the start time is in the correct format
			start, err := time.Parse("15:04", record[2])
			if err != nil {
				return fmt.Errorf("line %d: invalid start time format", i+1)
			}

			// Check if the end time is in the correct format
			end, err := time.Parse("15:04", record[3])
			if err != nil {
				return fmt.Errorf("line %d: invalid end time format", i+1)
			}

			// Check if the overtime is in the correct format and calculate the expected overtime
			overtime, err := strconv.Atoi(record[4])
			if err != nil {
				return fmt.Errorf("line %d: invalid overtime format", i+1)
			}
			minutes := int(end.Sub(start).Minutes()) - (config.DefaultWorkingHours * 60) - (config.BreakTime)

			// Check if the calculated overtime matches the expected overtime
			if minutes != overtime {
				return fmt.Errorf("line %d: overtime value is incorrect", i+1)
			}
		}
	}

	// No errors, return nil
	return nil
}
