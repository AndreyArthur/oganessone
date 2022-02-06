package main

import (
	"flag"
	"log"
	"net"

	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	google_grpc "google.golang.org/grpc"
)

func main() {
	var environment string
	flag.StringVar(&environment, "env", "test", "Specify the environment. Default is test.")
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
