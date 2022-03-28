package storage

import (
	"html/template"
	"oshino29/uraaka/post"
	"oshino29/uraaka/utils"
	"strings"
)

const ExpireDay float64 = 3

func (s *Storage) AddPost(p *post.Post) bool {
	_, err := s.db.Exec("INSERT INTO posts (post, html, time) VALUES (?, ?, ?)", p.Text, p.TextHTML, p.Time)
	return err == nil
}

func (s *Storage) AddRawPost(p *post.Post) bool {
	// when command /censor detected, add censorword separated by space to database and return
	if strings.HasPrefix(p.Text, "/censor"){
		words := strings.Split(p.Text, " ")
		for index := 1; index < len(words); index++ {
			if len(words[index]) == 0 { continue }
			s.AddCensorWord(words[index])
		}
		return true
	}

	p.RawToHtml(s.LoadCensorWords(), "#")
	return s.AddPost(p)
}

// should take time to rewrite these loadPosts related stuff
// for more reusability, like loadPosts(startDate, endDate)
func (s *Storage) LoadAllPosts() []post.Post {
	//query row
	rows, _ := s.db.Query("SELECT post, html, time FROM posts ORDER by time DESC")
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

func (s *Storage) LoadRecentPosts(days ...float64) []post.Post {
	// using default expire days when no arguent given
	var expireDay float64 = ExpireDay
	if len(days) != 0 {
		// or set to given argument
		expireDay = days[0]
	}

	//query row
	rows, _ := s.db.Query("SELECT post, html, time FROM posts ORDER by time DESC")
	defer rows.Close()

	ppp := make([]post.Post, 0)
	p := post.Post{}
	var html string
	for rows.Next() {
		rows.Scan(&p.Text, &html, &p.Time)
		p.TextHTML = template.HTML(html)
		// convert 2022-03-28 02:26:53 to 2022-03-28 02:26 for more compact view wehn display
		p.Time = utils.StripSecondsFromTime(p.Time)
		// p.TimeHTML = p.Time

		// only add post to memory when post date is between x days before to now
		if hoursPast := utils.HoursPast(p.Time); 0 <= hoursPast && hoursPast <= 24*expireDay {
			ppp = append(ppp, p)
		}
	}
	return ppp
}