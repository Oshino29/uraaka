package main

import (
	"html/template"
	"net/http"
	"oshino29/uraaka/post"
	"oshino29/uraaka/storage"
	"path"
	"time"
	"oshino29/uraaka/storage/migrate"
)

var DB string = "/data/uraaka.db"
var S *storage.Storage = storage.New(DB)

type PageData struct {
	// Pagetext string
	Posts post.Posts
}

func main() {
	http.HandleFunc("/", ShowPosts)
	http.HandleFunc("/newpost", NewPost)
	http.HandleFunc("/migrate", Migrate)
	http.HandleFunc("/censor", NewCensor)

	http.ListenAndServe("0.0.0.0:8080", nil)
}

func NewPost(rw http.ResponseWriter, r *http.Request) {
	p := post.Post{Text: (r.FormValue("body")), Time: time.Now().Format("2006-01-02 15:04:05")}

	S.AddRawPost(&p)
	http.Redirect(rw, r, r.Header.Get("Referer"), http.StatusFound)
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(fp))

	ppp := S.LoadPosts()
	data := PageData{
		Posts: ppp,
	}
	tmpl.Execute(w, data)
}

func Migrate(w http.ResponseWriter, r *http.Request) {
	m := migrate.New(DB, "", "")
	m.Migrate()
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func NewCensor(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}