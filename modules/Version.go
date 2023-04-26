package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var SoftwareVersion = 0.5

func GetVersion() {
	fmt.Printf("Version: %.2f\n", SoftwareVersion)
	fmt.Println("Checking for updates...")
	latestVersion, _ := getLatestVersion()
	fmt.Printf("Latest version: %.2f\n", latestVersion)
	if latestVersion > SoftwareVersion {
		fmt.Println("\n---------------------\nNew update available!")
	}
}

func getLatestVersion() (float64, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/FritzJo/ttracker/main/modules/Version.go")
	if err != nil {
		fmt.Println("Error: Cant fetch latest version.")
		return 0.0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
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
