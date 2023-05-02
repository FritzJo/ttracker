package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	m "example.com/ttracker/modules"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide an argument!")
	}
	recordFileName := strconv.Itoa(time.Now().Year()) + "_data.csv"
	recordList := m.ReadRecords(recordFileName)
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

	m.WriteRecords(recordFileName, recordList)
}
