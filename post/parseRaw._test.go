package post

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
func TestRawToHtml(t *testing.T) {
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
		{
			"cxcu\nifewf",
			"c◯◯u<br>ifewf",
		},
	}

	var p Post
	for _, data := range testCensorData {
		p.Text = data.raw
		p.RawToHtml(&censorList, "◯")

		t.Logf("\nraw: \t\t\t%s\nexpected: \t\t%s\nCensor() returned \t%s", data.raw, data.censored, string(p.TextHTML))
		if data.censored != string(p.TextHTML) {
			t.Errorf("the result doesn't behave as expected\n\n")
		} 
	}
}