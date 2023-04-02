package modules

import (
	m "example.com/ttracker/modules"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary JSON file with test data
	testConfig := `{
        "InitialOvertime": 60,
        "DefaultWorkingHours": 8
    }`
	tmpfile, err := ioutil.TempFile("", "testconfig*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}(tmpfile.Name()) // clean up
	if _, err := tmpfile.Write([]byte(testConfig)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Call LoadConfig with the temporary file path
	conf := m.LoadConfig(tmpfile.Name())

	// Verify that the returned Configuration has the expected values
	expectedInitialOvertime := 60
	if conf.InitialOvertime != expectedInitialOvertime {
		t.Errorf("Expected InitialOvertime to be %d, but got %d", expectedInitialOvertime, conf.InitialOvertime)
	}

	expectedDefaultWorkingHours := 8
	if conf.DefaultWorkingHours != expectedDefaultWorkingHours {
		t.Errorf("Expected DefaultWorkingHours to be %d, but got %d", expectedDefaultWorkingHours, conf.DefaultWorkingHours)
	}
}
