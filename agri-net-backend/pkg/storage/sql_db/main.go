package sql_db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

// NewStorage ...
func NewStorage(username, password, host, dbname string) (*sql.DB, error) {
	// Preparing the statement
	postgresStatment := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",host, username , password, dbname)
	db, err := sql.Open("postgres", postgresStatment)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
