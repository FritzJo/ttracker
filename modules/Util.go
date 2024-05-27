package modules

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/FritzJo/ttracker/modules/datatypes"
)

func PreviousYearRecordsExist() bool {
	var currentYear int = time.Now().Year()
	var fileName string = strconv.Itoa(currentYear-1) + "_data.csv"
	// Check if file for previous year exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func ReadRecords(recordFileName string) []datatypes.TimeRecord {
	// Read existing CSV data
	config, _ := datatypes.LoadConfig("config.json")
	fullPath := filepath.Join(config.StorageLocation, recordFileName)
	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
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
	var timeRecords []datatypes.TimeRecord
	for i, line := range data {
		if i > 0 {
			var record datatypes.TimeRecord
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

func WriteRecords(recordFileName string, recordList []datatypes.TimeRecord) {
	config, _ := datatypes.LoadConfig("config.json")
	// Write file for new records
	fw, err := os.Create(filepath.Join(config.StorageLocation, recordFileName))
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
	config, _ := datatypes.LoadConfig("config.json")
	startHour, _ := strconv.Atoi(strings.Split(workStart, ":")[0])
	startMinute, _ := strconv.Atoi(strings.Split(workStart, ":")[1])
	endHour, _ := strconv.Atoi(strings.Split(workEnd, ":")[0])
	endMinute, _ := strconv.Atoi(strings.Split(workEnd, ":")[1])
	return (endHour*60 + endMinute) - (startHour*60 + startMinute) - (config.DefaultWorkingHours * 60) - (config.BreakTime)
}

func LastRecordIsOpen(recordList []datatypes.TimeRecord) bool {
	if len(recordList) == 0 {
		return false
	}
	lastRecord := recordList[len(recordList)-1]
	return lastRecord.WorkEnd == ""
}
