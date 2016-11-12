package storage

import (
	"time"
)

type Table struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Key       string `gorm:"unique"`
	Val       string
	Table     string `gorm:"-"`
}

func (t *Table) TableName() string {
	return t.Table
}
