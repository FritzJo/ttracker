package datatypes

import (
	"time"
)

type TimeRecord struct {
	RecordType      string
	Date            time.Time
	WorkStart       string
	WorkEnd         string
	MinutesOvertime int
}
