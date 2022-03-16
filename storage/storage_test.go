package storage

import (
	// "database/sql"
	"testing"
	// "fmt"
	"strconv"
)

var db = "uraaka_test.db"

func (s *Storage) ListPosts(t *testing.T) {
	rows, _ := s.db.Query("SELECT id, post, time FROM posts")
	var id int
	var post string
	var time string
	for rows.Next() {
		rows.Scan(&id, &post, &time)
		// fmt.Println(strconv.Itoa(id) + ": " + time + "\n" + time)
		t.Logf("id: %s\ntime: %s\npost: %s", strconv.Itoa(id), time, post)
	}
}

func TestNew(t *testing.T) {
	s, err := New(db)
	if err != nil {
		t.Errorf("failed to create %s", db)
	}
	s.Init()
	s.ListPosts(t)
}

func TestInit(t *testing.T) {
	s, err := New(db)
	if err != nil {
		t.Errorf(err.Error())
	}

	statement, _ := s.db.Prepare("INSERT INTO posts (post, time) VALUES (?, ?)")
	statement.Exec("试试", "2022-03-16 21:40")

	s.ListPosts(t)
}