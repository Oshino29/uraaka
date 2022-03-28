package utils

import (
	"fmt"
	"time"
	"errors"
)

// var TimeFormats = map[string]string{
// 	"inFile": "2006-01-02_1504",
// 	"inFile_suffix": "2006-01-02_1504.post",
//     "inHtml": "2006-01-02 15:04",
// 	"inHtmlStash": "2006-01-02_15:04",
// }

var TimeFormats = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02_15:04:05",
	"2006-01-02_15:04",
	"2006-01-02_1504",
	"2006-01-02_150405",
	"2006-01-02_1504.post",
}

const LOCALE = "Asia/Shanghai"

// return the hours between of given string and now
func HoursPast(timeString string) float64 {
	// time.lo

	timeParsed, err := ParseTime(timeString)
	if err != nil {
		fmt.Print(err.Error())
		return -1
	}
	// return time.Since(timeParsed).Hours()
	fmt.Printf("time.Now() in utils/dateCalcu.go/HoursPast() has returned:\n\t%s\n", time.Now().String())
	return time.Now().Sub(timeParsed).Hours()
}

func ParseTime(timeString string) (time.Time, error) {
	for _, format := range TimeFormats {
        postTime, err := time.ParseInLocation(format, timeString, time.Local)
        if err == nil {
            return postTime, nil
        }
    }

	errString := "can't parse " + timeString + " to time.Time\n"
	return time.Time{}, errors.New(errString)
}