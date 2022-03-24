package storage

import (
	"database/sql"
	"log"
	// "fmt" 
	// "strconv"
	_ "github.com/mattn/go-sqlite3"
)


type Storage struct {
	db *sql.DB
}

func New(path string) *Storage {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("can't create or open database" + "\n" + err.Error())
	}

	db.SetMaxOpenConns(1)

	s := &Storage{ db: db}
	s.Init()
	// if err := Init(db); err != nil {
	// 	log.Panicf("can't init database" + "\n" + err.Error())
	// 	return nil, err
	// }
	
	return s

}
func (s *Storage) Init() error {
	sql := `	
			CREATE TABLE IF NOT EXISTS posts (
				id INTEGER PRIMARY KEY,
				post TEXT,
				html TEXT,
				time TEXT
			);

			CREATE TABLE IF NOT EXISTS censor_words (
				id INTERGER PRIMARY KEY,
				word TEXT
			);

			END;
	`
	_, err := s.db.Exec(sql)
	return err
}