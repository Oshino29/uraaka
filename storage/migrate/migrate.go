package migrate

import (
	"log"
	"oshino29/uraaka/storage"
)

// func (s *Storage) censorWordsFromFile(path string) error {
	
// }
type Migrate struct {
	Storage *storage.Storage
	postsPath, censorWordsPath string
}

//when empty string as parameter, default value will be used
//default dbPath: uraaka.db
// default postsPath: posts
// default cnesorWordsPath: censorList.txt
func New(dbPath, postsPath, censorWordsPath string) *Migrate {
	var m Migrate

	//  set to default value when not specified
	if dbPath == "" {dbPath = "uraaka.db"}
	if postsPath == "" {postsPath = "posts"}
	if censorWordsPath == "" {censorWordsPath = "censorList.txt"}

	var err error
	m.Storage, err = storage.New(dbPath)
	if err != nil {
		log.Fatalf("can't open %s while init Migrate.Storage", dbPath)
	}

	m.postsPath = postsPath
	m.censorWordsPath = censorWordsPath

	return &m
}

// CensorWords should be migrated before Posts
// cause Storage.AddRawPost calls 
func (m *Migrate) Migrate() {
	m.migrateCensorWords()
	m.migratePosts()
}