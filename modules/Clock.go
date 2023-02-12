package modules

import (
	"fmt"
	"time"
)

func clockIn() TimeRecord {
	var newRecord TimeRecord
	t := time.Now()
	newRecord.RecordType = "R"
	newRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	hours, minutes, _ := time.Now().Clock()
	newRecord.WorkStart = fmt.Sprintf("%d:%02d", hours, minutes)
	newRecord.WorkEnd = ""
	newRecord.MinutesOvertime = 0

	return newRecord
}

func clockOut(openRecord TimeRecord) TimeRecord {
	hours, minutes, _ := time.Now().Clock()
	openRecord.WorkEnd = fmt.Sprintf("%d:%02d", hours, minutes)
	openRecord.MinutesOvertime = CalcOvertime(openRecord.WorkStart, openRecord.WorkEnd)

	return openRecord
}
