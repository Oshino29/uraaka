package storage

import (
	"html/template"
	"oshino29/uraaka/post"
)

func (s *Storage) AddPost(p *post.Post) bool {
	_, err := s.db.Exec("INSERT INTO posts (post, html, time) VALUES (?, ?, ?)", p.Text, p.TextHTML, p.Time)
	return err == nil
}

func (s *Storage) AddRawPost(p *post.Post) bool {
	p.RawToHtml(s.LoadCensorWords(), "#")
	return s.AddPost(p)
}

func (s *Storage) LoadPosts() []post.Post {
	//query row
	rows, _ := s.db.Query("SELECT post, html, time FROM posts ORDER by id DESC")
	defer rows.Close()

	ppp := make([]post.Post, 0)
	p := post.Post{}
	var html string
	for rows.Next() {
		rows.Scan(&p.Text, &html, &p.Time)
		p.TextHTML = template.HTML(html)
		// p.TimeHTML = p.Time
		ppp = append(ppp, p)
	}
	return ppp
}
