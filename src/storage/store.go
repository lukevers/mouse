package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
	"time"
)

type Store struct {
	db    *gorm.DB
	t     *Table
	mutex *sync.Mutex
}

func New(dialect, dsn string) (*Store, error) {
	store := &Store{}
	var err error

	// Open a connection, and return an error if we failed to connect
	store.db, err = gorm.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	store.t = &Table{}
	store.mutex = &sync.Mutex{}

	return store, nil
}

func (s *Store) Table(table string) {
	s.t.Table = table
	if !s.db.HasTable(s.t) {
		s.db.CreateTable(s.t)
	}
}

func (s *Store) Get(table, key string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Table(table)

	var result Table
	s.db.Table(table).Where(Table{Key: key}).First(&result)

	return result.Val
}

func (s *Store) Put(table, key, val string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Table(table)

	var result Table
	s.db.Table(table).Where(Table{Key: key}).First(&result)

	if result.ID == 0 {
		s.db.Table(table).Create(&Table{
			Key:       key,
			Val:       val,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	} else {
		s.db.Table(table).Model(&result).Updates(Table{
			Val:       val,
			UpdatedAt: time.Now(),
		})
	}
}

func (s *Store) Delete(table, key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Table(table)

	var result Table
	s.db.Table(table).Where(Table{Key: key}).First(&result)

	if result.ID != 0 {
		s.db.Table(table).Delete(&result)
	}
}
