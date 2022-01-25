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
		goerr := godotenv.Load(".env.test")
		if goerr != nil {
			log.Fatal(goerr)
			return
		}
	}
	if environment == "production" {
		goerr := godotenv.Load(".env")
		if goerr != nil {
			log.Fatal(goerr)
			return
		}
	}
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	sql, err := db.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	migrator, err := database.NewMigrator(sql)
	if err != nil {
		log.Fatal(err)
		return
	}
	migrator.Down()
}
