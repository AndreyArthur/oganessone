package main

import (
	"flag"
	"log"

	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/joho/godotenv"
)

func main() {
	var environment string
	flag.StringVar(&environment, "e", "test", "Specify an environment. Default is test.")
	flag.Parse()
	if environment == "test" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	if environment == "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	migrator, err := database.NewMigrator(db.Connect())
	if err != nil {
		log.Fatal(err)
		return
	}
	migrator.Down()
}
