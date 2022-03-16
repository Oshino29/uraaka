package utils

import (
	"testing"
)
var censorList []string = []string{
	"xc", "北京", "xc", "bgm.tv",
}

type censorData struct {
	raw 		string
	censored 	string
}
func TestCensor(t *testing.T) {
	testCensorData := []censorData{
		{
			"xc found x y x c in cx",
			"◯◯ found x y x c in cx",
		},
		{
			"我在北京吃了顿饭啊",
			"我在◯◯吃了顿饭啊",
		},
		{
			"cxcuifewf",
			"c◯◯uifewf",
		},
	}

	for _, data := range testCensorData {
		result := Censor(&data.raw, &censorList, "◯")

		t.Logf("\nraw: \t\t\t%s\nexpected: \t\t%s\nCensor() returned \t%s", data.raw, data.censored, *result)
		if data.censored != *result {
			t.Errorf("the result doesn't behave as expected\n\n")
		} 
	}
}

 