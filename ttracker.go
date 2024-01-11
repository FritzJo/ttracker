package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	m "github.com/FritzJo/ttracker/modules"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide an argument!")
	}
	recordFileName := strconv.Itoa(time.Now().Year()) + "_data.csv"
	recordList := m.ReadRecords(recordFileName)

	// Import previous year
	if m.PreviousYearRecordsExist() {
		conf := m.LoadConfig("config.json")
		recordFileName := strconv.Itoa(time.Now().Year()-1) + "_data.csv"
		oldRecordList := m.ReadRecords(recordFileName)
		currentOvertimeAmount := 0
		for _, record := range oldRecordList {
			currentOvertimeAmount += record.MinutesOvertime
		}
		if conf.InitialOvertime != currentOvertimeAmount {
			fmt.Println("First new record for this year, summarizing last years records...")
			fmt.Println("Last years overtime: " + strconv.Itoa(currentOvertimeAmount))
			conf.InitialOvertime = currentOvertimeAmount
			m.SaveConfig("config.json", conf)
		}
	}
	// Check for empty file
	if len(recordList) >= 1 || os.Args[1] == "in" {
		switch os.Args[1] {
		case "in":
			recordList = m.In(recordList, os.Args)
		case "out":
			recordList = m.Out(recordList, os.Args)
		case "status":
			fmt.Println(m.Status(recordList))
		case "summary":
			recordList = m.Summary(recordList)
		case "take":
			recordList = m.Take(recordList)
		case "update":
			m.GetVersion()
		case "validate":
			err := m.ValidateCSVFile(recordFileName)
			if err == nil {
				fmt.Println("OK!")
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("Not enough data available for this command. Please clock in at least once.")
	}

	m.WriteRecords(recordFileName, recordList)
}
