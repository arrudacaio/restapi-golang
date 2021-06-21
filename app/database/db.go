package database

import (
	"github.com/jmoiron/sqlx"
)

type PostDB interface {
	Open() error
	Close() error
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	return nil
}

func (d *DB) Close() error {
	return nil
}
