package storage

// import (
// 	"testing"
// 	"oshino29/uraaka/post"
// )

// func TestSavePost(t *testing.T) {
// 	p := &post.Post{"测试写入", "", "2022-03-17 22:00", ""}

// 	s, err := New(db)
// 	if err != nil {
// 		t.Errorf("failed to create *Storage")
// 	}

// 	s.AddPost(p)

// 	s.ListPosts(t)
// 	// //query row
// 	// rows, _ := s.db.Query("SELECT id, post, time FROM posts")
// 	// var id int
// 	// var post string
// 	// var time string
// 	// for rows.Next() {
// 	// 	rows.Scan(&id, &post, &time)
// 	// 	// fmt.Println(strconv.Itoa(id) + ": " + time + "\n" + time)
// 	// 	t.Logf("id: %s\ntime: %s\npost: %s", strconv.Itoa(id), time, post)
// 	// }
// }

// func TestLoadPosts(t *testing.T) {
// 	s, err := New(db)
// 	if err != nil {
// 		t.Errorf("failed to create data.db")
// 	}

// 	ppp := s.LoadPosts()
// 	// ppp, _ := TempLoadPosts()

// 	for _, p := range ppp {
// 		t.Logf("\ntime: %s\n%s\n\n", p.Time, p.Text)
// 	}
// }
