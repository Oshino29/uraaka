package utils

import (
	"testing"
)

type stripData struct {
	input string
	expect string
}

func TestStripSecondsFromTime(t *testing.T) {
	testData := []stripData{
		{
			input: "2022-03-28 02:26:53",
			expect: "2022-03-28 02:26",
		},
		{
			input: "2022-03-27 17:55",
			expect: "2022-03-27 17:55",
		},
	}

	for _, td := range testData {
		if result := StripSecondsFromTime(td.input); result != td.expect {
			t.Errorf("input:\t%s\nexpected:\t%s\noutput:\t%s\n", td.input, td.expect, result)
		}
	}
}