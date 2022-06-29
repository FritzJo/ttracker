package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type TimeRecord struct {
	Date            time.Time
	WorkStart       string
	WorkEnd         string
	MinutesOvertime int
}

func readRecords(data [][]string) []TimeRecord {
	var timeRecords []TimeRecord
	for i, line := range data {
		if i > 0 {
			var record TimeRecord
			fmt.Println(line)
			record.Date, _ = time.Parse("2006-01-02", line[0])
			record.WorkStart = line[1]
			record.WorkEnd = line[2]
			record.MinutesOvertime, _ = strconv.Atoi(line[3])
			timeRecords = append(timeRecords, record)
		}
	}
	return timeRecords
}

func clockIn() TimeRecord {
	var newRecord TimeRecord
	t := time.Now()
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
	openRecord.MinutesOvertime = calcOvertime(openRecord.WorkStart, openRecord.WorkEnd)

	return openRecord
}

func calcOvertime(workStart string, workEnd string) int {
	startHour, _ := strconv.Atoi(strings.Split(workStart, ":")[0])
	startMinute, _ := strconv.Atoi(strings.Split(workStart, ":")[1])
	endHour, _ := strconv.Atoi(strings.Split(workEnd, ":")[0])
	endMinute, _ := strconv.Atoi(strings.Split(workEnd, ":")[1])

	return (endHour*60 + endMinute) - (startHour*60 + startMinute) - (7 * 60) - (60)

}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide an argument!")
	}

	// Read existing CSV data
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Parse and print records
	recordList := readRecords(data)
	fmt.Println("\n")
	fmt.Printf("%+v\n", recordList)
	fmt.Println(len(recordList))

	var rec TimeRecord
	switch os.Args[1] {
	case "in":
		rec = clockIn()
	case "out":
		rec = clockOut(rec)
	}
	fmt.Println(rec)
	//fmt.Println(calcOvertime("7:55","16:15"))
}
