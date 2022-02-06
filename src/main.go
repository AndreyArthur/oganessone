package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/joho/godotenv"
	google_grpc "google.golang.org/grpc"
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
	flag.StringVar(&environment, "env", "test", "Specify the environment. Default is test.")
	flag.Parse()
	if environment == "production" {
		envFile := strings.Join([]string{dirname, "/../.env"}, "")
		goerr := godotenv.Load(envFile)
		if goerr != nil {
			log.Fatal(goerr)
			return
		}
	} else {
		envFile := strings.Join([]string{dirname, "/../.env.test"}, "")
		goerr := godotenv.Load(envFile)
		if goerr != nil {
			log.Fatal(goerr)
			return
		}
	}
	googleGrpcServer := google_grpc.NewServer()
	server, err := grpc.NewGrpcServer(googleGrpcServer)
	if err != nil {
		log.Fatal(err)
		return
	}
	listener, goerr := net.Listen("tcp", "0.0.0.0:50051")
	if goerr != nil {
		log.Fatal(err)
		return
	}
	server.Start(listener)
}
