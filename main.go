package main

import (
	"html/template"
	"net/http"
	"path"
	"time"
	// "github.com/russross/blackfriday"
)

type Post struct {
	Text []byte
	Time string
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
	http.ListenAndServe(":8080", nil)

}

func HandlePost(rw http.ResponseWriter, r *http.Request) {
	p := Post{Text: []byte(r.FormValue("body")), Time: time.Now().Format("2006-01-02_1504")}
	// SavePost(&post)
	p.Save()
	rw.Write(p.Text)
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(fp))

	data := PageData{
		// PageText: "My TODO list",
		Posts: []Post{
			{Text: []byte("Task 1"), Time: "2022-01-03"},
			{Text: []byte("Task 1"), Time: "2022-02-09"},
			{Text: []byte("Task 1"), Time: "2022-02-20"},
		},
	}
	tmpl.Execute(w, data)
}
