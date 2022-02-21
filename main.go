package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"time"
	// "github.com/russross/blackfriday"
)

type Post struct {
	Text string
	TextHTML template.HTML
	Time string
	TimeHTML string
}
type Posts []Post

type PageData struct {
	// Pagetext string
	Posts Posts
}

func main() {
	http.HandleFunc("/newpost", HandlePost)
	// http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", ShowPosts)
	http.ListenAndServe("0.0.0.0:8080", nil)

}

func HandlePost(rw http.ResponseWriter, r *http.Request) {
	p := Post{Text: (r.FormValue("body")), Time: time.Now().Format("2006-01-02_1504")}
	// SavePost(&post)
	p.Save()
	rw.Write([]byte(p.Text))
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(fp))

	// var ppp Posts
	ppp := make([]Post, 0)

	// 读取 posts/ 文件夹下面所有 filename
	filenames, err := os.ReadDir("posts")
	if err != nil {
		log.Fatal(err)
	}

	for i := len(filenames) -1 ; i >= 0; i--{
		ppp = append(ppp, *LoadPost(filenames[i].Name()))
	}
	// ppp = append(ppp, *LoadPost("2022-02-21_1837"))
	// ppp = []Post{
	// 	*LoadPost("2022-02-21_1837"),
	// 	*LoadPost("2022-02-21_1455"),
	// 	*LoadPost("2022-02-21_1453"),
	// }
	data := PageData{
		// PageText: "My TODO list",
		Posts: ppp,
	}
	tmpl.Execute(w, data)
}
