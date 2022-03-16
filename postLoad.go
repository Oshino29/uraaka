package main

// import (
// 	"errors"
// 	"html/template"
// 	"io/ioutil"
// 	"log"
// 	"strings"
// 	"time"
// 	// "github.com/russross/blackfriday"
// )

// func (ppp *Posts) Load() {
//     ppp = append(*ppp, Post{Text: "", Time: "2022-99-99"})
// }
// const (
//     layoutFILE = "2006-01-02_1504"
// 	layoutFILE_Suffix = "2006-01-02_1504.post"
//     layoutHTML  = "2006-01-02 15:04"
// )
var TimeFormats = map[string]string{
	"inFile": "2006-01-02_1504",
	"inFile_suffix": "2006-01-02_1504.post",
    "inHtml": "2006-01-02 15:04",
}

// func LoadPost(postID string) *Post {
// 	// if postTime has ".post" suffix, remove it
// 	// postTime = strings.TrimSuffix(postTime, ".post")
// 	// //get the full filename with path and suffix
// 	// filename := "posts/" + postTime + ".post"
// 	postTime, err  := parsePostID(postID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	filename := "posts/" + postTime.Format(TimeFormats["inFile_suffix"])

// 	//read file content
// 	b, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	safe := string(b)
// 	//safe = template.HTMLEscapeString(safe)
// 	safe = strings.Replace(safe, "\n", "<br>", -1)

// 	//create a pointer to this Post struct, in which all loaded info is stored
	
// 	p := &Post{
// 		Text:     "",
// 		TextHTML: template.HTML(censor(safe)),
// 		Time:     "",
// 	}
// 	return p
// }

// func parsePostID(postID string) (time.Time, error) {
// 	for _, format := range TimeFormats {
//         postTime, err := time.Parse(format, postID)
//         if err == nil {
//             return postTime, nil
//         }
//     }
// 	return time.Time{}, errors.New("can't parse postID to time.Time type")
// }