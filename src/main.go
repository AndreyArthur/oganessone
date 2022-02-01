package main

import (
	"log"
	"net"

	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc"
	google_grpc "google.golang.org/grpc"
)

func main() {
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
