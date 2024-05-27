package datatypes

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	InitialOvertime     int
	DefaultWorkingHours int
	BreakTime           int
	StorageLocation     string
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

func SaveConfig(configPath string, conf Configuration) error {
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error, couldn't create configuration file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(conf)
	if err != nil {
		return fmt.Errorf("error, couldn't encode configuration: %v", err)
	}

	return nil
}
