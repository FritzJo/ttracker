package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ValidateCSVFile(filepath string) error {
	config := LoadConfig("config.json")

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 5
	csvReader.Comma = ';'
	csvReader.TrimLeadingSpace = true

	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	for i, record := range records {
		if i != 0 {
			if len(record) != csvReader.FieldsPerRecord {
				return fmt.Errorf("line %d: invalid number of fields", i+1)
			}

			_, err := time.Parse("2006-01-02", record[1])
			if err != nil {
				return fmt.Errorf("line %d: invalid date format", i+1)
			}

			start, err := time.Parse("15:04", record[2])
			if err != nil {
				return fmt.Errorf("line %d: invalid start time format", i+1)
			}

			end, err := time.Parse("15:04", record[3])
			if err != nil {
				return fmt.Errorf("line %d: invalid end time format", i+1)
			}

			overtime, err := strconv.Atoi(record[4])
			if err != nil {
				return fmt.Errorf("line %d: invalid overtime format", i+1)
			}

			minutes := int(end.Sub(start).Minutes()) - (config.DefaultWorkingHours * 60) - (config.BreakTime)
			if minutes != overtime {
				return fmt.Errorf("line %d: overtime value is incorrect", i+1)
			}
		}

	}

	return nil
}
