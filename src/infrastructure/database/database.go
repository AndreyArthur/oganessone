package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct{}

func (database *Database) Connect() *sql.DB {
	postgres, err := NewPostgres()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	info := postgres.GenerateConnectionString()
	connection, err := sql.Open("postgres", info)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func NewDatabase() (*Database, error) {
	return &Database{}, nil
}
