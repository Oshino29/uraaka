package utils

import (
	"strings"
	"unicode/utf8"
)

func Censor(rawText *string, censorList *[]string, masks ...string) *string {
	// masks 是用来替代敏感词的字符，默认为 #
	masks = append(masks, "#")

	// 把 censorList 中的替换词以 ("北京", "##") 的形式添加进 oldnew array
	var oldnew []string
	for _, censorWord := range *censorList {
		oldnew = append(oldnew, censorWord, strings.Repeat(masks[0], utf8.RuneCountInString(censorWord)))
	}

	//用 string.Replacer 来返回 censor 过的 string
	//... 用来将 oldnew flatten 成 oldnew 中的元素们
	r := strings.NewReplacer(oldnew...)
	censored := r.Replace(*rawText)
	return &censored
}