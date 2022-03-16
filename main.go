package main

import (
	"html/template"
	"net/http"
	"oshino29/uraaka/post"
	"oshino29/uraaka/storage"
	"path"
	"time"
)

// type Post struct {
// 	Text     string
// 	TextHTML template.HTML
// 	Time     string
// 	TimeHTML string
// }
// type Posts []Post

var DB string = "uraaka.db"

type PageData struct {
	// Pagetext string
	Posts post.Posts
}

func main() {
	http.HandleFunc("/newpost", NewPost)
	// http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", ShowPosts)
	http.ListenAndServe("0.0.0.0:8080", nil)

}

func NewPost(rw http.ResponseWriter, r *http.Request) {
	p := post.Post{Text: (r.FormValue("body")), Time: time.Now().Format(TimeFormats["inHtml"])}

	s, err := storage.New(DB)
	if err != nil {
		return
	}

	s.AddRawPost(&p)
	http.Redirect(rw, r, r.Header.Get("Referer"), http.StatusFound)
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(fp))

	s, err := storage.New(DB)
	if err != nil {
		return
	}

	ppp := s.LoadPosts()
	data := PageData{
		Posts: ppp,
	}
	tmpl.Execute(w, data)
}
