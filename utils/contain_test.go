package utils

import (
	"testing"
)

type containData struct {
	x, y 	*[]string
	result 	bool
}
func TestContain(t *testing.T) {
	testContainData := []containData{
		{
			&[]string{"xc", "北京", "bgm.tv"},
			&[]string{"xc", "北京", "bgm.tv"},
			true,
		},
		{
			&[]string{"xc", "北京", "xc", "bgm.tv",},
			&[]string{"xc", "北京", "bgm.tv"},
			true,
		},
		{
			&[]string{"xc", "北京"},
			&[]string{"xc", "北京", "bgm.tv"},
			false,
		},
	}

	for _, data := range testContainData {
		t.Logf("check if slice %v contains slice %v", data.x, data.y)
		t.Logf("expected result: \t\t%t", data.result)
		if Contain(data.x, data.y) != data.result {
			t.Errorf("result from Contain(): \t%t\nnot identical", !data.result)
		} else {
			t.Logf("result from Contain(): \t%t", data.result)
		}
	}
}
    