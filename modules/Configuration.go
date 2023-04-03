package modules

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	InitialOvertime     int
	DefaultWorkingHours int
	BreakTime           int
}

func LoadConfig(configPath string) Configuration {
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Error, couldn't open configuration file: ", err)
		fmt.Println(configPath)
		return Configuration{}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	Conf := Configuration{}
	err = decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Error, couldn't decode configuration: ", err)
	}
	return Conf
}
