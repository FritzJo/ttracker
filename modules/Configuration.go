package modules

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	InitialOvertime     int
	DefaultWorkingHours int
}

func LoadConfig(configPath string) Configuration {
	file, _ := os.Open(configPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	Conf := Configuration{}
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Error, couldn't read configuration: ", err)
	}
	return Conf
}
