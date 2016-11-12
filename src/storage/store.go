package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Store struct {
	db *gorm.DB
}

func New(dialect, dsn string) (*Store, error) {
	store := &Store{}
	var err error

	// Open a connection, and return an error if we failed to connect
	store.db, err = gorm.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	return store, nil
}
