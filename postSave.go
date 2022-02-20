package main

import (
	"log"
	"os"
)

// func SavePost(p *[]byte) {
// 	now := time.Now().Format("2006-01-02_1504")
// 	filename := "posts/" + now + ".post"

// 	err := os.WriteFile(filename, *p, 0644)
// 	if err != nil {
// 	    log.Fatal(err)
// 	}
// }

func (p *Post) Save() {
	filename := "posts/" + p.Time + ".post"

	err := os.WriteFile(filename, p.Text, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
