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
	RecordType      string
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
			record.RecordType = line[0]
			record.Date, _ = time.Parse("2006-01-02", line[1])
			record.WorkStart = line[2]
			record.WorkEnd = line[3]
			record.MinutesOvertime, _ = strconv.Atoi(line[4])
			timeRecords = append(timeRecords, record)
		}
	}
	return timeRecords
}

func clockIn() TimeRecord {
	var newRecord TimeRecord
	t := time.Now()
	newRecord.RecordType = "R"
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

	recordFileName := strconv.Itoa(time.Now().Year()) + "_data.csv"
	// Read existing CSV data
	f, err := os.OpenFile(recordFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

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
	fmt.Println("Reading existing records")
	recordList := readRecords(data)
	fmt.Println()

	// Write file for new records
	fw, err := os.Create(recordFileName)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(fw)
	csvWriter.Comma = ';'
	defer csvWriter.Flush()

	var rec TimeRecord
	switch os.Args[1] {
	case "in":
		fmt.Println("Clocking in")
		rec = clockIn()
		recordList = append(recordList, rec)
	case "out":
		fmt.Println("Clocking out")
		lastRecord := recordList[len(recordList)-1]
		if lastRecord.WorkEnd == "" {
			// TODO: This doesnt check for record type R yet!
			rec = clockOut(lastRecord)
			recordList = recordList[:len(recordList)-1]
			recordList = append(recordList, rec)
		} else {
			fmt.Println("Can't clock out, because there is currently no open time record!")
		}
	case "summary":
		fmt.Println("Creating summary")
		currentOvertimeAmount := 0
		for _, record := range recordList {
			fmt.Println(record.MinutesOvertime)
			currentOvertimeAmount += record.MinutesOvertime
		}
		fmt.Println("\n=> " + strconv.Itoa(currentOvertimeAmount))
	case "take":
		fmt.Println("Taking time off: " + os.Args[2])
		var offRecord TimeRecord
		offRecord.RecordType = "O"
		offtime, _ := strconv.Atoi(os.Args[2])
		offRecord.MinutesOvertime = -1 * offtime
		recordList = append(recordList, offRecord)
	}

	// Write csv headers
	err = csvWriter.Write([]string{"type", "date", "start", "end", "overtime"})
	if err != nil {
		log.Fatal(err)
	}

	// Writing data to csv
	for _, record := range recordList {
		e := csvWriter.Write([]string{
			record.RecordType,
			record.Date.Format("2006-01-02"),
			record.WorkStart,
			record.WorkEnd,
			strconv.Itoa(record.MinutesOvertime)})
		if e != nil {
			fmt.Println(e)
		}
		if err := csvWriter.Error(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done")
}
