package database

import (
	"fmt"
	"os"
)

type Postgres struct{}

func (postgres *Postgres) GenerateConnectionString() string {
	host, port, user, password, name :=
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME")
	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)
	return info
}

func NewPostgres() (*Postgres, error) {
	return &Postgres{}, nil
}
