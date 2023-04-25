package modules

import (
	"testing"
	"time"

	m "example.com/ttracker/modules"
)

func TestClockIn(t *testing.T) {
	tr := m.ClockIn()
	// Verify the returned TimeRecord has the correct RecordType
	if tr.RecordType != "R" {
		t.Errorf("Expected RecordType to be 'R', but got '%s'", tr.RecordType)
	}
	// Verify the returned TimeRecord has the current date
	expectedDate := time.Now().Format("2006-01-02")
	if tr.Date.Format("2006-01-02") != expectedDate {
		t.Errorf("Expected Date to be '%s', but got '%s'", expectedDate, tr.Date.Format("2006-01-02"))
	}
	// Verify the returned TimeRecord has the correct WorkStart time
	expectedStart := time.Now().Format("15:04")
	if tr.WorkStart != expectedStart {
		t.Errorf("Expected WorkStart to be '%s', but got '%s'", expectedStart, tr.WorkStart)
	}
	// Verify the returned TimeRecord has an empty WorkEnd field
	if tr.WorkEnd != "" {
		t.Errorf("Expected WorkEnd to be empty, but got '%s'", tr.WorkEnd)
	}
	// Verify the returned TimeRecord has 0 MinutesOvertime
	if tr.MinutesOvertime != 0 {
		t.Errorf("Expected MinutesOvertime to be 0, but got '%d'", tr.MinutesOvertime)
	}
}
func TestClockInWithOptionalParameter(t *testing.T) {
	// Clock in with a specific work start time
	workStart := "10:00"
	record := m.ClockIn(workStart)
	// Check that the record has the correct values
	if record.RecordType != "R" {
		t.Errorf("Expected RecordType to be 'R', but got '%s'", record.RecordType)
	}
	expectedDate := time.Now().Format("2006-01-02")
	if record.Date.Format("2006-01-02") != expectedDate {
		t.Errorf("Expected Date to be '%s', but got '%s'", expectedDate, record.Date.Format("2006-01-02"))
	}
	if record.WorkStart != workStart {
		t.Errorf("Expected WorkStart to be '%s', but got '%s'", workStart, record.WorkStart)
	}
	if record.WorkEnd != "" {
		t.Errorf("Expected WorkEnd to be empty, but got '%s'", record.WorkEnd)
	}
	if record.MinutesOvertime != 0 {
		t.Errorf("Expected MinutesOvertime to be 0, but got %d", record.MinutesOvertime)
	}
}
func TestClockOut(t *testing.T) {
	// Create a TimeRecord to use as input for ClockOut
	inputRecord := m.TimeRecord{
		RecordType: "R",
		Date:       time.Now(),
		WorkStart:  "09:00",
		WorkEnd:    "",
	}
	// Call ClockOut and get the output TimeRecord
	outputRecord := m.ClockOut(inputRecord)
	// Verify that the outputRecord has the expected WorkEnd value
	expectedEnd := time.Now().Format("15:04")
	if outputRecord.WorkEnd != expectedEnd {
		t.Errorf("Expected WorkEnd to be '%s', but got '%s'", expectedEnd, outputRecord.WorkEnd)
	}
	// Verify that the outputRecord has the expected MinutesOvertime value
	expectedOvertime := m.CalcOvertime(inputRecord.WorkStart, outputRecord.WorkEnd)
	if outputRecord.MinutesOvertime != expectedOvertime {
		t.Errorf("Expected MinutesOvertime to be %d, but got %d", expectedOvertime, outputRecord.MinutesOvertime)
	}
}
func TestClockOutWithOptionalParameter(t *testing.T) {
	// Create a time record with a specific work start time
	workStart := "10:00"
	record := m.TimeRecord{
		RecordType:      "R",
		Date:            time.Now(),
		WorkStart:       workStart,
		WorkEnd:         "",
		MinutesOvertime: 0,
	}
	// Clock out with a specific work end time
	workEnd := "17:00"
	record = m.ClockOut(record, workEnd)
	// Check that the record has the correct values
	if record.RecordType != "R" {
		t.Errorf("Expected RecordType to be 'R', but got '%s'", record.RecordType)
	}
	expectedDate := time.Now().Format("2006-01-02")
	if record.Date.Format("2006-01-02") != expectedDate {
		t.Errorf("Expected Date to be '%s', but got '%s'", expectedDate, record.Date.Format("2006-01-02"))
	}
	if record.WorkStart != workStart {
		t.Errorf("Expected WorkStart to be '%s', but got '%s'", workStart, record.WorkStart)
	}
	if record.WorkEnd != workEnd {
		t.Errorf("Expected WorkEnd to be '%s', but got '%s'", workEnd, record.WorkEnd)
	}
}
