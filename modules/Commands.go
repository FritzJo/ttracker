package modules

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/FritzJo/ttracker/modules/datatypes"
)

// In function is used to clock in a new time record with the current time.
// If an argument is provided, it will be used as the work start time for the new record.
func In(recordList []datatypes.TimeRecord, args []string) []datatypes.TimeRecord {
	message := ""
	if len(args) > 2 {
		rec := ClockIn(args[2])
		recordList = append(recordList, rec)
		message = fmt.Sprintf("Clocking in (%s)", args[2])
	} else {
		rec := ClockIn()
		recordList = append(recordList, rec)
		message = fmt.Sprintf("Clocking in (%s)", time.Now().Format("15:04"))
	}
	fmt.Println(message)
	return recordList
}

// Out clocks the user out and adds a new time record to the input recordList.
//
// Parameters:
// - recordList: A list of TimeRecord structs representing the user's past working time times.
// - args: A slice of strings containing the command-line arguments passed to the program.
//
// Returns:
// The input recordList, with a new TimeRecord appended if the user was clocked out successfully.
func Out(recordList []datatypes.TimeRecord, args []string) []datatypes.TimeRecord {
	// Print a message indicating that the user is clocking out
	message := ""

	// Get the last time record from the list
	lastRecord := recordList[len(recordList)-1]

	// If the last time record has an empty WorkEnd field, remove it from the list and create a new time record with the
	// current time as the WorkEnd field
	if lastRecord.WorkEnd == "" {
		// TODO: This doesn't check for record type R yet!
		recordList = recordList[:len(recordList)-1]
		if len(args) > 2 {
			rec := ClockOut(lastRecord, args[2])
			recordList = append(recordList, rec)
			message = fmt.Sprintf("Clocking out (%s)", args[2])
		} else {
			rec := ClockOut(lastRecord)
			recordList = append(recordList, rec)
			message = fmt.Sprintf("Clocking out (%s)", time.Now().Format("15:04"))
		}

		fmt.Println(message)
		// Print the number of overtime minutes worked during the day
		fmt.Println("Today's overtime: " + strconv.Itoa(recordList[len(recordList)-1].MinutesOvertime))
	} else {
		// If the last time record already has a non-empty WorkEnd field, print an error message
		fmt.Println("Can't clock out, because there is currently no open time record!")
	}

	// Return the input recordList, with a new TimeRecord appended if the user was clocked out successfully
	return recordList
}

// Summary prints a summary of the user's overtime hours based on the given list of time records.
//
// Parameters:
// - recordList: A list of TimeRecord structs representing the user's past working time times.
//
// Returns:
// The input recordList, unchanged.
func Summary(recordList []datatypes.TimeRecord) []datatypes.TimeRecord {
	// Print a message indicating that a summary is being created
	fmt.Println("Creating summary")

	// Load the initial overtime amount from the configuration file and print it
	currentOvertimeAmount := datatypes.LoadConfig("config.json").InitialOvertime
	fmt.Println("Initial overtime: " + strconv.Itoa(currentOvertimeAmount) + " min")

	// Iterate over the list of time records and print each record's date and overtime minutes
	for _, record := range recordList {
		minutes := fmt.Sprintf("%4d", record.MinutesOvertime)
		fmt.Println(record.Date.Format("2006-01-02") + " -> " + minutes + " min")
		currentOvertimeAmount += record.MinutesOvertime
	}

	// Print the total overtime hours accumulated
	fmt.Println("\n=> " + strconv.Itoa(currentOvertimeAmount) + " min")

	// Return the input recordList, unchanged
	return recordList
}

// Take adds a new TimeRecord to the given list indicating that the user is taking time off.
//
// Parameters:
// - recordList: A list of TimeRecord structs representing the user's past working time times.
//
// Returns:
// A new list of TimeRecord structs that includes the new TimeRecord representing the time off.
func Take(recordList []datatypes.TimeRecord) []datatypes.TimeRecord {
	// Print a message indicating that the user is taking time off
	fmt.Println("Taking time off: " + os.Args[2])

	// Create a new TimeRecord for the time off
	var offRecord datatypes.TimeRecord
	offRecord.RecordType = "O"
	t := time.Now().Local()
	offRecord.Date, _ = time.Parse("2006-01-02", t.Format("2006-01-02"))
	offtime, _ := strconv.Atoi(os.Args[2])
	offRecord.MinutesOvertime = -1 * offtime

	// Add the new TimeRecord to the list and return the modified list
	recordList = append(recordList, offRecord)
	return recordList
}

// Status returns a string containing information about the user's current work status.
//
// Parameters:
// - recordList: A list of TimeRecord structs representing the user's past working time times.
//
// Returns:
// A formatted string containing the clock-in time and overtime minutes.
func Status(recordList []datatypes.TimeRecord) string {
	// Get the most recent time record from the list
	openRecord := recordList[len(recordList)-1]

	// Get the current time and calculate the overtime minutes
	hours, minutes, _ := time.Now().Clock()
	currentTime := fmt.Sprintf("%02d:%02d", hours, minutes)
	openRecord.MinutesOvertime = CalcOvertime(openRecord.WorkStart, currentTime)

	// Return a formatted string containing the clock-in time and overtime minutes
	return fmt.Sprintf("Clocked in at: %v\nOvertime: %d Minutes.",
		openRecord.WorkStart,
		openRecord.MinutesOvertime)
}
