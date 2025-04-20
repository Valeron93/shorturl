package data

import "database/sql"

type UrlEntry struct {
	Id  int    `json:"id,omitempty"`
	Key string `json:"key"`
	Url string `json:"url"`
}

type UrlEntryTable struct {
	db *sql.DB
}

func NewUrlEntryTable(db *sql.DB) (UrlEntryTable, error) {

	query := `
	CREATE TABLE IF NOT EXISTS UrlEntries (
    	id INTEGER PRIMARY KEY,
    	key TEXT,
		url TEXT
	)
	`
	_, err := db.Exec(query)
	if err != nil {
		return UrlEntryTable{}, err
	}

	return UrlEntryTable{
		db: db,
	}, nil
}

func (u *UrlEntryTable) CreateNewUrlEntry(entry UrlEntry) error {

	query := `
	INSERT INTO UrlEntries(key, url) VALUES (?, ?)
	`
	_, err := u.db.Exec(query, entry.Key, entry.Url)
	return err
}

func (u *UrlEntryTable) GetUrlByKey(key string) (UrlEntry, error) {
	var entry UrlEntry
	query := `
	SELECT id, key, url FROM UrlEntries WHERE key = ?
	`
	err := u.db.QueryRow(query, key).Scan(&entry.Id, &entry.Key, &entry.Url)
	return entry, err
}
