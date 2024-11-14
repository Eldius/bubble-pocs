package storm

import (
	"fmt"
	"github.com/asdine/storm/v3"
)

var (
	db *storm.DB
)

func newDB(s string) (*storm.DB, error) {
	db, err := storm.Open(s)
	if err != nil {
		err = fmt.Errorf("opening db file for '%s': %w", s, err)
		return nil, err
	}
	return db, nil
}

// GetDB returns a DB pool
func GetDB(s string) (*storm.DB, error) {
	if db == nil {
		var err error
		db, err = newDB(s)
		if err != nil {
			err = fmt.Errorf("opening db file for '%s': %w", s, err)
			return nil, err
		}
	}
	return db, nil
}
