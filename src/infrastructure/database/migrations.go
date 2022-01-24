package database

import (
	"database/sql"
	"log"

	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type Migrator struct {
	db *sql.DB
}

func (migrator *Migrator) Up() {
	db := migrator.db
	defer db.Close()
	_, err := db.Query("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = db.Query(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
			username VARCHAR(16) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(60) UNIQUE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (migrator *Migrator) Down() {
	db := migrator.db
	defer db.Close()
	_, err := db.Query("DROP TABLE IF EXISTS users;")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = db.Query("DROP EXTENSION IF EXISTS \"uuid-ossp\";")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewMigrator(db *sql.DB) (*Migrator, *shared.Error) {
	return &Migrator{
		db: db,
	}, nil
}
