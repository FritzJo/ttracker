package modules

import (
	m "example.com/ttracker/modules"
	"testing"
	"time"
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
