package utils

import (
	"testing"
	"time"
)
type dateCalcuData struct {
	timeString string
	expectHours float64
}

func TestHoursPast(t *testing.T) {
	errData := []string{
		"2006-01-02_15:4",
	}


	for _, testDate := range errData {
		if result := HoursPast(testDate); result != -1 {
			t.Errorf("function HoursPast returned %f (expect -1) when invalid string %s was passed", result, testDate)
		}
	}

	testData := []string{
		"2022-03-24 13:09",
		"2006-01-02 15:04",
		"2022-03-02 15:04",
		"2022-03-26 00:04",
		"2022-03-28 01:30:00",
	}

	loc, _ := time.LoadLocation("")
	t.Logf("current timezone: %s\n", loc)
	for _, testDate := range testData {
		result := HoursPast(testDate)
		t.Logf("function HoursPast returned <%f> when string <%s> was passed", result, testDate)
		if result == -1 {
			t.Errorf("function HoursPast returned unexpected -1 when valid string %s was passed", testDate)
		}
	}

	t.Logf("current timezone settings returned by time package:")
	tzLoc := time.Now().Location().String()
    t.Logf(tzLoc)
    t.Logf(time.Now().String())
}

func TestParseTime(t *testing.T) {
	testData := []string{
		"2022-03-24 13:09",
		"2006-01-02 15:04",
		"2022-03-02 15:04",
		"2022-03-26 00:04",
		"2022-03-28 01:30:00",
	}

	for _, testString := range testData {
		resultTime, _ := ParseTime(testString)
		t.Logf("------")
		t.Logf("%s\t%s\t%s", "string", "parsed result", "locale in result")
		t.Logf("%s\t%s\t%s", testString, resultTime.String(), resultTime.Location().String())
	}

	t.Logf(time.Local.String())
}