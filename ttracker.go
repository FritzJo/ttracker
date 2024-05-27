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

	// Import previous year
	if modules.PreviousYearRecordsExist() {
		conf := datatypes.LoadConfig("config.json")
		recordFileName := strconv.Itoa(time.Now().Year()-1) + "_data.csv"
		oldRecordList := modules.ReadRecords(recordFileName)
		currentOvertimeAmount := 0
		for _, record := range oldRecordList {
			currentOvertimeAmount += record.MinutesOvertime
		}
		if conf.InitialOvertime != currentOvertimeAmount {
			fmt.Println("First new record for this year, summarizing last years records...")
			fmt.Println("Last years overtime: " + strconv.Itoa(currentOvertimeAmount))
			conf.InitialOvertime = currentOvertimeAmount
			datatypes.SaveConfig("config.json", conf)
		}
	}
	// Check for empty file
	if len(recordList) >= 1 || os.Args[1] == "in" {
		switch os.Args[1] {
		case "in":
			recordList = modules.In(recordList, os.Args)
		case "out":
			recordList = modules.Out(recordList, os.Args)
		case "status":
			fmt.Println(modules.Status(recordList))
		case "summary":
			recordList = modules.Summary(recordList)
		case "take":
			recordList = modules.Take(recordList)
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
