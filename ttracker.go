package main

import (
	m "example.com/ttracker/modules"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide an argument!")
	}

	recordFileName := strconv.Itoa(time.Now().Year()) + "_data.csv"

	// Parse and print records
	fmt.Println("Reading existing records")
	recordList := m.ReadRecords(recordFileName)
	fmt.Println()

	switch os.Args[1] {
	case "in":
		recordList = m.In(recordList)
	case "out":
		recordList = m.Out(recordList)
	case "summary":
		recordList = m.Summary(recordList)
	case "take":
		recordList = m.Take(recordList)
	}

	m.WriteRecords(recordFileName, recordList)
}
