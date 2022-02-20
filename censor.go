package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
	"log"
)

var censorList []string
var CensorListFile = "censorList.txt"

// 把储存的 censorList 文件对的每行存进数组
func loadCensorList() {
	// 用 ioutil 把文件存入 memory
	fileBytes, err := ioutil.ReadFile(CensorListFile)
	// 如果无法打开文件，报错
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 以 \n 为分隔符，将 memory 中的数据存入数组
	censorList = strings.Split(string(fileBytes), "\n")
}

func censor(rawText string) string {
	// 加载 censorList
	loadCensorList()
	// 把 censorList 中的替换词以 ("北京", "##") 的形式添加进 oldnew array
	var oldnew []string
	for _, censorWord := range censorList {
		oldnew = append(oldnew, censorWord, strings.Repeat("#", utf8.RuneCountInString(censorWord)))
	}

	//用 string.Replacer 来返回 censor 过的 string
	//... 用来将 oldnew flatten 成 oldnew 中的元素们
	r := strings.NewReplacer(oldnew...)
	censored := r.Replace(rawText)
	return censored
}

func appendCensorWord(censorWord string) {
	//Append second line
	file, err := os.OpenFile(CensorListFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("second line"); err != nil {
		log.Fatal(err)
	}
}
