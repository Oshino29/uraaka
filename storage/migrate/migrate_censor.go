package migrate

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	// "unicode/utf8"
	// "log"
)

var censorList []string
var CensorListFile = "censorList.txt"

// 这个函数原本是在 plainfile 版本中用来把 censorList.txt 中的 censorWords 存进 memory
// 现在用来把 censorList 中的 censorWords 存进 database，和 censor_test.go 一起食用来 run this single function
func (m *Migrate) migrateCensorWords() error {
	// 用 ioutil 把文件存入 memory
	fileBytes, err := ioutil.ReadFile(m.censorWordsPath)
	// 如果无法打开文件，报错
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 以 \n 为分隔符，将 memory 中的数据存入数组
	censorList = strings.Split(string(fileBytes), "\n")
	
	
	// 打开 database，将 memory 中的 censorWords 存入
	// s, err := New(DB)
	// if err != nil {
	// 	return err
	// }
	for _, w := range censorList {
		m.Storage.AddCensorWord(w)
	}
	// censorList = *s.LoadCensorWords()
	return err
}