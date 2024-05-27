package datatypes

import (
	"strconv"
	"time"
)

type TimeRecord struct {
	RecordType      string
	Date            time.Time
	WorkStart       string
	WorkEnd         string
	MinutesOvertime int
}

func CreateNewTimeRecord(inputTime string) TimeRecord {
	var newRecord TimeRecord
	t := time.Now().Local()
	newRecord.RecordType = "R"
	newRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	newRecord.WorkStart = inputTime

	newRecord.WorkEnd = ""
	newRecord.MinutesOvertime = 0
	return newRecord
}

func CreateNewOffRecord(inputTime string) TimeRecord {
	var offRecord TimeRecord
	offRecord.RecordType = "O"
	t := time.Now().Local()
	offRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	offtime, _ := strconv.Atoi(inputTime)
	offRecord.MinutesOvertime = -1 * offtime
	return offRecord
}
