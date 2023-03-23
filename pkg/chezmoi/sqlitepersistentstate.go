package chezmoi

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SQLitePersistentState struct {
	db *sql.DB
}

func NewSQLitePersistentState(dataSourceName string) (*SQLitePersistentState, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &SQLitePersistentState{
		db: db,
	}, nil
}

func (s *SQLitePersistentState) Close() error {
	return s.db.Close()
}

func (s *SQLitePersistentState) CopyTo(_ PersistentState) error {
	return nil // FIXME
}

func (s *SQLitePersistentState) Data() (any, error) {
	return nil, nil // FIXME
}

func (s *SQLitePersistentState) Delete(bucket, key []byte) error {
	return nil // FIXME
}

func (s *SQLitePersistentState) DeleteBucket(bucket []byte) error {
	return nil // FIXME
}

func (s *SQLitePersistentState) ForEach(bucket []byte, fn func([]byte, []byte) error) error {
	return nil // FIXME
}

func (s *SQLitePersistentState) Get(bucket, key []byte) ([]byte, error) {
	return nil, nil // FIXME
}

func (s *SQLitePersistentState) Set(bucket, key, value []byte) error {
	return nil // FIXME
}
