package modules

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadRecords(recordFileName string) []TimeRecord {
	// Read existing CSV data
	f, err := os.OpenFile(recordFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var timeRecords []TimeRecord
	for i, line := range data {
		if i > 0 {
			var record TimeRecord
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

func WriteRecords(recordFileName string, recordList []TimeRecord) {
	// Write file for new records
	fw, err := os.Create(recordFileName)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(fw)
	csvWriter.Comma = ';'
	defer csvWriter.Flush()
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
}

// CalcOvertime calculates the overtime in minutes based on the given work start and end times.
// It takes the work start and end times as strings in the format "hh:mm", and uses the default working
// hours and break time from the config file to calculate the overtime.
// It returns the overtime in minutes as an integer.
func CalcOvertime(workStart string, workEnd string) int {
	config := LoadConfig("config.json")
	startHour, _ := strconv.Atoi(strings.Split(workStart, ":")[0])
	startMinute, _ := strconv.Atoi(strings.Split(workStart, ":")[1])
	endHour, _ := strconv.Atoi(strings.Split(workEnd, ":")[0])
	endMinute, _ := strconv.Atoi(strings.Split(workEnd, ":")[1])
	return (endHour*60 + endMinute) - (startHour*60 + startMinute) - (config.DefaultWorkingHours * 60) - (config.BreakTime)
}
