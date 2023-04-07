package modules

import (
	"fmt"
	"time"
)

func ClockIn(workStart ...string) TimeRecord {
	var newRecord TimeRecord
	t := time.Now()
	newRecord.RecordType = "R"
	newRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	if len(workStart) > 0 && workStart[0] != "" {
		newRecord.WorkStart = workStart[0]
	} else {
		hours, minutes, _ := time.Now().Clock()
		newRecord.WorkStart = fmt.Sprintf("%d:%02d", hours, minutes)
	}

	newRecord.WorkEnd = ""
	newRecord.MinutesOvertime = 0

	return newRecord
}

func ClockOut(openRecord TimeRecord, workEnd ...string) TimeRecord {
	if len(workEnd) > 0 && workEnd[0] != "" {
		openRecord.WorkEnd = workEnd[0]
	} else {
		hours, minutes, _ := time.Now().Clock()
		openRecord.WorkEnd = fmt.Sprintf("%d:%02d", hours, minutes)
	}
	openRecord.MinutesOvertime = CalcOvertime(openRecord.WorkStart, openRecord.WorkEnd)

	return openRecord
}
