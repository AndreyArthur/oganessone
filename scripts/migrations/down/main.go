package main

import (
	"flag"
	"log"

	"github.com/AndreyArthur/oganessone/src/infrastructure/factories"
)

func main() {
	var environment string
	flag.StringVar(&environment, "e", "test", "Specify an environment. Default is test.")
	flag.Parse()
	migrator, err := factories.MakeMigrator(environment)
	if err != nil {
		log.Fatal(err)
		return
	}
	migrator.Down()
}
