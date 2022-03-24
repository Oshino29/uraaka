package storage

import (
	"testing"
	"oshino29/uraaka/utils"
)

func TestAddCensorWord(t *testing.T) {
	s := New("censor_test.db")

	s.AddCensorWord("CensorTestWord")

	rows, _ := s.db.Query("SELECT word FROM censor_words WHERE word = 'CensorTestWord'")
	defer rows.Close()

	var found bool = false
	var w string
	for rows.Next() {
		rows.Scan(&w)
		if w == "CensorTestWord" {found = true}
	}
	if found == false {
		t.Errorf("CensroTestWord not added to test database")
	} else {
		t.Logf("found the added CensroTestWord in test database")
	}
		

}

func TestLoadCensorWords(t *testing.T) {
	s := New("censor_test.db")

	test_www := &([]string{"xc", "北京", "bgm.tv"})
	for _, test_w := range *test_www {
		s.AddCensorWord(test_w)
	}
	// s.AddCensorWord("xc")
	// s.AddCensorWord("北京")
	// s.AddCensorWord("zhihu.com")
	www := s.LoadCensorWords()

	if len(*www) == 0 {
		t.Errorf("*Storage.LoadCensorWords returned an empty slice")
	}
	
	t.Logf("\ntest slice: %v\nloaded slice: %v\n", test_www, www)
	if utils.Contain(www, test_www) {
		t.Logf("all test words are contained in the slice returned by *Storage.LoadCensorWords")

	} else {
		t.Errorf("slice returned by *Storage.LoadCensorWords does not contain all test words")
	}
}