package modules

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	m "github.com/FritzJo/ttracker/modules"
)

func TestReadRecords(t *testing.T) {
	// Create a temporary file and write some sample data to it
	f, err := ioutil.TempFile("", "test_records.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	sampleData := [][]string{
		{"type", "date", "start", "end", "overtime"},
		{"R", "2022-01-01", "09:00", "17:00", "0"},
		{"R", "2022-01-02", "09:00", "16:00", "-60"},
	}
	csvWriter := csv.NewWriter(f)
	csvWriter.Comma = ';'
	for _, line := range sampleData {
		if err := csvWriter.Write(line); err != nil {
			t.Fatal(err)
		}
	}
	csvWriter.Flush()
	// Call ReadRecords on the temporary file
	records := m.ReadRecords(f.Name())
	// Verify the records are as expected
	if len(records) != 2 {
		t.Errorf("Expected 2 records, got %d", len(records))
	}
	expectedRecord1 := m.TimeRecord{
		RecordType:      "R",
		Date:            time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		WorkStart:       "09:00",
		WorkEnd:         "17:00",
		MinutesOvertime: 0,
	}
	if !reflect.DeepEqual(records[0], expectedRecord1) {
		t.Errorf("Expected record 1 to be %+v, got %+v", expectedRecord1, records[0])
	}
	expectedRecord2 := m.TimeRecord{
		RecordType:      "R",
		Date:            time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
		WorkStart:       "09:00",
		WorkEnd:         "16:00",
		MinutesOvertime: -60,
	}
	if !reflect.DeepEqual(records[1], expectedRecord2) {
		t.Errorf("Expected record 2 to be %+v, got %+v", expectedRecord2, records[1])
	}
}
func TestWriteRecords(t *testing.T) {
	// Create a temporary file to write to
	f, err := ioutil.TempFile("", "test_records.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	// Create some sample records to write
	records := []m.TimeRecord{
		{
			RecordType:      "R",
			Date:            time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			WorkStart:       "09:00",
			WorkEnd:         "17:00",
			MinutesOvertime: 0,
		},
		{
			RecordType:      "R",
			Date:            time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			WorkStart:       "09:00",
			WorkEnd:         "16:00",
			MinutesOvertime: -60,
		},
	}
	// Call WriteRecords to write the records to the file
	m.WriteRecords(f.Name(), records)
	// Read the file and verify that it contains the expected data
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	expectedData := "type;date;start;end;overtime\nR;2022-01-01;09:00;17:00;0\nR;2022-01-02;09:00;16:00;-60\n"
	if string(data) != expectedData {
		t.Errorf("Expected file data to be %q, got %q", expectedData, string(data))
	}
}
func TestCalcOvertime(t *testing.T) {
	// Define some test cases
	testCases := []struct {
		workStart        string
		workEnd          string
		expectedOvertime int
	}{
		{"09:00", "17:00", -60},
		{"08:30", "17:30", 0},
		{"07:00", "19:00", 180},
	}
	// Iterate over the test cases and verify the results
	for _, tc := range testCases {
		overtime := m.CalcOvertime(tc.workStart, tc.workEnd)
		if overtime != tc.expectedOvertime {
			t.Errorf("CalcOvertime(%q, %q) = %d, expected %d", tc.workStart, tc.workEnd, overtime, tc.expectedOvertime)
		}
	}
}
