package modules

import (
	"fmt"
	"os"
	"strconv"
)

func In(recordList []TimeRecord) []TimeRecord {
	fmt.Println("Clocking in")
	rec := clockIn()
	recordList = append(recordList, rec)
	return recordList
}

func Out(recordList []TimeRecord) []TimeRecord {
	fmt.Println("Clocking out")
	lastRecord := recordList[len(recordList)-1]
	if lastRecord.WorkEnd == "" {
		// TODO: This doesnt check for record type R yet!
		rec := clockOut(lastRecord)
		recordList = recordList[:len(recordList)-1]
		recordList = append(recordList, rec)
	} else {
		fmt.Println("Can't clock out, because there is currently no open time record!")
	}

	return recordList
}


func Summary(recordList []TimeRecord) []TimeRecord {
	fmt.Println("Creating summary")
	currentOvertimeAmount := 0
	for _, record := range recordList {
		fmt.Println(record.MinutesOvertime)
		currentOvertimeAmount += record.MinutesOvertime
	}
	fmt.Println("\n=> " + strconv.Itoa(currentOvertimeAmount))
	return recordList
}

func Take(recordList []TimeRecord) []TimeRecord {
	fmt.Println("Taking time off: " + os.Args[2])
	var offRecord TimeRecord
	offRecord.RecordType = "O"
	offtime, _ := strconv.Atoi(os.Args[2])
	offRecord.MinutesOvertime = -1 * offtime
	recordList = append(recordList, offRecord)
	return recordList
}