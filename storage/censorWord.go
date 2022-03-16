package storage

func (s *Storage) AddCensorWord(w string) bool {
	_, err := s.db.Exec("INSERT INTO censor_words (word) VALUES (?)", w)
	return err == nil
}
func (s *Storage) LoadCensorWords() *[]string {
	// rows, _ := s.db.Query("SELECT word FROM censor_words ORDER by LEN(word) DESC")
	rows, _ := s.db.Query("SELECT word FROM censor_words")
	defer rows.Close()

	var w string
	www := make([]string, 0)
	for rows.Next() {
		rows.Scan(&w)
		www = append(www, w)
	}
	return &www
}