package datatypes

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Configuration struct {
	InitialOvertime     int
	DefaultWorkingHours int
	BreakTime           int
	StorageLocation     string
}

var (
	config *Configuration
	once   sync.Once
)

func LoadConfig(configPath string) (*Configuration, error) {
	var erro error
	once.Do(func() {
		file, err := os.Open(configPath)
		if err != nil {
			erro = err // Capture the error so it can be returned
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		config = &Configuration{} // Initialize the package-level config
		erro = decoder.Decode(config)
	})

	if erro != nil {
		return nil, fmt.Errorf("error loading config: %w", erro)
	}
	return config, nil
}
