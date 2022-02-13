package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AndreyArthur/oganessone/src/infrastructure/factories"
)

func main() {
	var environment string
	var up, down bool
	flag.StringVar(&environment, "env", "test", "Specify an environment. Default is test.")
	flag.BoolVar(&up, "up", false, "Specify if you want to migrate up or down.")
	flag.BoolVar(&down, "down", false, "Specify if you want to migrate up or down.")
	flag.Parse()
	if !up && !down {
		fmt.Println("Cannot migrate. You must specify -up or -down.")
		return
	}
	migrator, err := factories.MakeMigrator(environment)
	if err != nil {
		log.Fatal(err)
		return
	}
	if up {
		migrator.Up()
		fmt.Println("Migrated up successfully.")
		return
	}
	if down {
		migrator.Down()
		fmt.Println("Migrated down successfully.")
		return
	}
}
