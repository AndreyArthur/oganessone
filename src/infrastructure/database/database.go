package database

import (
	"database/sql"
	"log"

	"github.com/AndreyArthur/oganessone/src/core/shared"
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
	connection, goerr := sql.Open("postgres", info)
	if goerr != nil {
		log.Fatal(goerr)
		return nil
	}
	return connection
}

func NewDatabase() (*Database, *shared.Error) {
	return &Database{}, nil
}
