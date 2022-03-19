package migrate

import (
	"errors"
	// "html/template"
	"io/ioutil"
	"log"
	// "strings"
	"time"
	"oshino29/uraaka/post"
	"os"
)

// func (ppp *post.Posts) Load() {
//     ppp = append(*ppp, post.Postost{Text: "", Time: "2022-99-99"})
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

func (m *Migrate) migratePosts() {
	filenames, err := os.ReadDir(m.postsPath)
	if err != nil {
		log.Fatal(err)
	}

	// iterate through filanames in reversed order
	// ppp := make([]post.Post, 0)
	for i := len(filenames) - 1; i >= 0; i-- {
		m.Storage.AddRawPost(LoadPostFromFile(filenames[i].Name()))
	}
}

func LoadPostFromFile(postID string) *post.Post {
	// if postTime has ".post" suffix, remove it
	// postTime = strings.TrimSuffix(postTime, ".post")
	// //get the full filename with path and suffix
	// filename := "posts/" + postTime + ".post"
	postTime, err  := parsePostID(postID)
	if err != nil {
		log.Fatal(err)
	}
	filename := "posts/" + postTime.Format(TimeFormats["inFile_suffix"])

	//read file content
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	t := string(b)
	//safe = template.HTMLEscapeString(safe)
	// safe = strings.Replace(safe, "\n", "<br>", -1)

	//create a pointer to this Post struct, in which all loaded info is stored
	
	p := &post.Post{
		Text:     t,
		Time:     postTime.Format(TimeFormats["inHtml"]),
	}
	return p
}

func parsePostID(postID string) (time.Time, error) {
	for _, format := range TimeFormats {
        postTime, err := time.Parse(format, postID)
        if err == nil {
            return postTime, nil
        }
    }
	return time.Time{}, errors.New("can't parse postID to time.Time type")
}