package main

import (
	"flag"
	"log"

	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
)

func main() {
	var environment string
	flag.StringVar(&environment, "e", "test", "Specify an environment. Default is test.")
	flag.Parse()
	env, err := helpers.NewEnv()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = env.Load(environment)
	if err != nil {
		log.Fatal(err)
		return
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
	migrator.Up()
}
