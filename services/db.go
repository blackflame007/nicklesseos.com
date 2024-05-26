package service

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DatabaseService struct {
	db *sql.DB
}

func NewDatabaseService(dbUrl string) (*DatabaseService, error) {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, err
	}

	return &DatabaseService{db: db}, nil
}

func (ds *DatabaseService) Close() error {
	return ds.db.Close()
}
