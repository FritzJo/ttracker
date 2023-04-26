package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var SoftwareVersion = 1.0

func GetVersion() {
	fmt.Printf("Version: %.2f\n", SoftwareVersion)
	fmt.Println("Checking for updates...")
	updateAvailable, _ := checkForUpdate()

	if updateAvailable {
		fmt.Println("New version available!")
	}
}

func checkForUpdate() (bool, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/FritzJo/ttracker/main/modules/Version.go")
	if err != nil {
		fmt.Println("Error: Cant fetch latest version.")
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Cant fetch latest version.")
		return false, err
	}
	latestVersion := string(body)
	fmt.Println(latestVersion)
	//if latestVersion != SoftwareVersion {
	//	return true, nil
	//}
	return false, nil
}
