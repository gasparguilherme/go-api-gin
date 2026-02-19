package db

import (
	"database/sql"
	"go-api/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConnectionBase)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
