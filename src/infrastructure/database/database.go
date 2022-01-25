package database

import (
	"database/sql"
	"log"

	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	_ "github.com/lib/pq"
)

type Database struct{}

func (database *Database) Connect() (*sql.DB, *shared.Error) {
	postgres, err := NewPostgres()
	if err != nil {
		return nil, err
	}
	info := postgres.GenerateConnectionString()
	connection, goerr := sql.Open("postgres", info)
	if goerr != nil {
		log.Println(goerr)
		return nil, exceptions.NewInternalServerError()
	}
	return connection, nil
}

func NewDatabase() (*Database, *shared.Error) {
	return &Database{}, nil
}
