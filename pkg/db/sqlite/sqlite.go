package sqlite

import (
	"database/sql"
	//go-sqlite3 is important
	_ "github.com/mattn/go-sqlite3"
)

type Options struct {
	Address string
}

func Init(opts *Options) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", opts.Address)
	if err != nil {
		return nil, err
	}
	return db, nil
}
