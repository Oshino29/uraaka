package post

import (
	"html/template"
	"strings"
	"oshino29/uraaka/utils"
)

func (p *Post) RawToHtml(censorList *[]string, mask string) {
	s := template.HTMLEscapeString(p.Text)
	s = strings.Replace(s, "\n", "<br>", -1)

	s = *utils.Censor(&s, censorList, mask)
	p.TextHTML = template.HTML(s)
}