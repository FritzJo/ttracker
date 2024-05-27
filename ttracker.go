package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/FritzJo/ttracker/modules"
	"github.com/FritzJo/ttracker/modules/datatypes"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide an argument!")
	}
	recordFileName := strconv.Itoa(time.Now().Year()) + "_data.csv"
	recordList := modules.ReadRecords(recordFileName)
	currentTime := modules.GetCurrentTime()
	if len(os.Args) > 2 {
		currentTime = os.Args[2]
	}
	// Check for empty file
	if len(recordList) >= 1 || os.Args[1] == "in" {
		switch os.Args[1] {
		case "in":
			recordList = In(recordList, currentTime)
		case "out":
			recordList = Out(recordList, currentTime)
		case "status":
			Status(recordList)
		case "summary":
			Summary(recordList)
		case "take":
			recordList = Take(recordList, currentTime) // "currentTime" contains amount of time taken in this case
		}
	} else {
		fmt.Println("Not enough data available for this command. Please clock in at least once.")
	}

	modules.WriteRecords(recordFileName, recordList)
}

func In(recordList []datatypes.TimeRecord, currentTime string) []datatypes.TimeRecord {
	rec := datatypes.CreateNewTimeRecord(currentTime)
	fmt.Printf("Clocking in at %s\n", rec.WorkStart)
	recordList = append(recordList, rec)
	return recordList
}

func Out(recordList []datatypes.TimeRecord, currentTime string) []datatypes.TimeRecord {
	if modules.LastRecordIsOpen(recordList) {
		lastIndex := len(recordList) - 1
		// TODO: This doesn't check for record type R yet!
		recordList[lastIndex].WorkEnd = currentTime
		recordList[lastIndex].MinutesOvertime = modules.CalcOvertime(recordList[lastIndex].WorkStart, recordList[lastIndex].WorkEnd)
		fmt.Printf("Clocking out at %s\n", recordList[lastIndex].WorkEnd)
		fmt.Println("Today's overtime: " + strconv.Itoa(recordList[lastIndex].MinutesOvertime))
	} else {
		fmt.Println("Can't clock out, because there is currently no open time record!")
	}

	return recordList
}

func Summary(recordList []datatypes.TimeRecord) {
	fmt.Println("Creating summary")

	config, _ := datatypes.LoadConfig("config.json")
	currentOvertimeAmount := config.InitialOvertime
	fmt.Println("Initial overtime: " + strconv.Itoa(currentOvertimeAmount) + " min")

	for _, record := range recordList {
		minutes := fmt.Sprintf("%4d", record.MinutesOvertime)
		fmt.Println(record.Date.Format("2006-01-02") + " -> " + minutes + " min")
		currentOvertimeAmount += record.MinutesOvertime
	}

	fmt.Println("\n=> " + strconv.Itoa(currentOvertimeAmount) + " min")
}

func Take(recordList []datatypes.TimeRecord, timeOff string) []datatypes.TimeRecord {
	fmt.Println("Taking time off: " + timeOff)
	rec := datatypes.CreateNewOffRecord(timeOff)
	recordList = append(recordList, rec)
	return recordList
}

func Status(recordList []datatypes.TimeRecord) {
	openRecord := recordList[len(recordList)-1]
	openRecord.MinutesOvertime = modules.CalcOvertime(openRecord.WorkStart, modules.GetCurrentTime())

	fmt.Printf("Clocked in at: %v\nOvertime: %d Minutes.\n",
		openRecord.WorkStart,
		openRecord.MinutesOvertime)
}
