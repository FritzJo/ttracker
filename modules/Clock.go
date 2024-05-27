package modules

import (
	"fmt"
	"time"

	"github.com/FritzJo/ttracker/modules/datatypes"
)

// ClockIn creates a new TimeRecord with a record type of "R", a date of the current date, and a work start time
// of either the current time or the provided work start time (if present). The work end time is empty and the
// minutes of overtime is initialized to 0. Returns the new TimeRecord.
func ClockIn(workStart ...string) datatypes.TimeRecord {
	var newRecord datatypes.TimeRecord
	t := time.Now().Local()
	newRecord.RecordType = "R"
	newRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	if len(workStart) > 0 && workStart[0] != "" {
		newRecord.WorkStart = workStart[0]
	} else {
		hours, minutes, _ := time.Now().Clock()
		newRecord.WorkStart = fmt.Sprintf("%02d:%02d", hours, minutes)
	}
	newRecord.WorkEnd = ""
	newRecord.MinutesOvertime = 0
	return newRecord
}
