package main

import (
	"testing"
	m "example.com/ttracker/modules"
)

func TestCalcOvertime1(t *testing.T) {
	workStart := "00:00"
	workEnd := "08:00"
	expectedOutput := 0
	output := m.CalcOvertime(workStart, workEnd)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %c", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestCalcOvertime2(t *testing.T) {
	workStart := "01:00"
	workEnd := "08:00"
	expectedOutput := -60
	output := m.CalcOvertime(workStart, workEnd)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %c", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}

func TestCalcOvertime3(t *testing.T) {
	workStart := "0:00"
	workEnd := "00:45"
	expectedOutput := -435
	output := m.CalcOvertime(workStart, workEnd)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %c", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}
