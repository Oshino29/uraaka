package main

import (
	"net/http"

	// "github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/newpost", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)

}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	post := []byte(r.FormValue("body"))
	SavePost(&post)
	// markdown := blackfriday.MarkdownCommon(post)
	rw.Write(post)
}
