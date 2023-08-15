package modules

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var SoftwareVersion = 0.5

// GetVersion prints the current software version and checks for updates.
// It also prints the latest version available and informs the user if a new update is available.
func GetVersion() {
	// Print the current software version.
	fmt.Printf("Version: %.2f\n", SoftwareVersion)

	// Check for updates and print the latest version available.
	fmt.Println("Checking for updates...")
	latestVersion, _ := getLatestVersion()
	fmt.Printf("Latest version: %.2f\n", latestVersion)

	// Inform the user if a new update is available.
	if latestVersion > SoftwareVersion {
		fmt.Println("\n---------------------\nNew update available!")
	}
}

// getLatestVersion retrieves the latest version of the software from a GitHub repository.
// It sends an HTTP GET request to a specific URL and parses the response to extract the version number.
// The function returns the latest version number as a float64 and an error if any occurs.
func getLatestVersion() (float64, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/FritzJo/ttracker/main/modules/Version.go")
	if err != nil {
		fmt.Println("Error: Cant fetch latest version.")
		return 0.0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Cant fetch latest version.")
		return 0.0, err
	}
	latestVersion := 0.0
	for _, element := range strings.Split(string(body), "\n") {
		if strings.HasPrefix(element, "var SoftwareVersion = ") {
			if s, err := strconv.ParseFloat(strings.Split(element, "var SoftwareVersion = ")[1], 64); err == nil {
				latestVersion = s
			}
		}
	}
	return latestVersion, nil
}
