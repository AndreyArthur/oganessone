package main

import (
	"flag"
	"log"
	"strings"

	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/joho/godotenv"
)

func main() {
	path, err := helpers.NewPath()
	if err != nil {
		log.Fatal(err)
		return
	}
	filename, err := path.File()
	if err != nil {
		log.Fatal(err)
		return
	}
	dirname, err := path.Dir(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	var environment string
	flag.StringVar(&environment, "e", "test", "Specify an environment. Default is test.")
	flag.Parse()
	if environment == "production" {
		envFile := strings.Join([]string{dirname, "/../../../env"}, "")
		goerr := godotenv.Load(envFile)
		if goerr != nil {
			log.Fatal(goerr)
			return
		}
	} else {
		envFile := strings.Join([]string{dirname, "/../../../.env.test"}, "")
		goerr := godotenv.Load(envFile)
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
