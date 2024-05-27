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

	// Check for empty file
	if len(recordList) >= 1 || os.Args[1] == "in" {
		switch os.Args[1] {
		case "in":
			recordList = In(recordList, os.Args)
		case "out":
			recordList = Out(recordList, os.Args)
		case "status":
			result := Status(recordList)
			fmt.Println(result)
		case "summary":
			recordList = Summary(recordList)
		case "take":
			recordList = Take(recordList)
		case "validate":
			err := modules.ValidateCSVFile(recordFileName)
			if err == nil {
				fmt.Println("OK!")
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("Not enough data available for this command. Please clock in at least once.")
	}

	modules.WriteRecords(recordFileName, recordList)
}

func In(recordList []datatypes.TimeRecord, args []string) []datatypes.TimeRecord {
	var rec datatypes.TimeRecord

	if len(args) > 2 {
		rec = modules.ClockIn(args[2])
	} else {
		rec = modules.ClockIn()
	}
	fmt.Printf("Clocking in at %s\n", rec.WorkStart)
	recordList = append(recordList, rec)
	return recordList
}

func Out(recordList []datatypes.TimeRecord, args []string) []datatypes.TimeRecord {
	if modules.LastRecordIsOpen(recordList) {
		lastIndex := len(recordList) - 1
		// TODO: This doesn't check for record type R yet!
		if len(args) > 2 {
			recordList[lastIndex].WorkEnd = args[2]
		} else {
			recordList[lastIndex].WorkEnd = time.Now().Format("15:04")
		}
		fmt.Printf("Clocking out at %s\n", recordList[lastIndex].WorkEnd)
		fmt.Println("Today's overtime: " + strconv.Itoa(recordList[lastIndex].MinutesOvertime))
	} else {
		fmt.Println("Can't clock out, because there is currently no open time record!")
	}

	return recordList
}

func Summary(recordList []datatypes.TimeRecord) []datatypes.TimeRecord {
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
	return recordList
}

func Take(recordList []datatypes.TimeRecord) []datatypes.TimeRecord {
	fmt.Println("Taking time off: " + os.Args[2])

	var offRecord datatypes.TimeRecord
	offRecord.RecordType = "O"
	t := time.Now().Local()
	offRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	offtime, _ := strconv.Atoi(os.Args[2])
	offRecord.MinutesOvertime = -1 * offtime

	recordList = append(recordList, offRecord)
	return recordList
}

func Status(recordList []datatypes.TimeRecord) string {
	openRecord := recordList[len(recordList)-1]

	hours, minutes, _ := time.Now().Clock()
	currentTime := fmt.Sprintf("%02d:%02d", hours, minutes)
	openRecord.MinutesOvertime = modules.CalcOvertime(openRecord.WorkStart, currentTime)

	return fmt.Sprintf("Clocked in at: %v\nOvertime: %d Minutes.",
		openRecord.WorkStart,
		openRecord.MinutesOvertime)
}
