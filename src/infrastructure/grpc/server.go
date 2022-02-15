package grpc

import (
	"context"
	"log"
	"net"

	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/factories"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedAccountsServiceServer
}

func (*server) CreateAccount(
	ctx context.Context, request *protobuf.CreateAccountRequest,
) (*protobuf.CreateAccountResponse, error) {
	rpc, err := factories.MakeCreateAccountRpc()
	if err != nil {
		return &protobuf.CreateAccountResponse{
			Error: &protobuf.Error{
				Type:    err.Type,
				Name:    err.Name,
				Message: err.Message,
			},
			Data: nil,
		}, nil
	}
	return rpc.Perform(ctx, request)
}

type GrpcServer struct {
	googleGrpcServer *grpc.Server
	protoServer      protobuf.AccountsServiceServer
}

func (gs *GrpcServer) Start(lis net.Listener) {
	protobuf.RegisterAccountsServiceServer(gs.googleGrpcServer, gs.protoServer)
	err := gs.googleGrpcServer.Serve(lis)
	gs.googleGrpcServer.Stop()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewGrpcServer(googleGrpcServer *grpc.Server) (*GrpcServer, *shared.Error) {
	server := &server{}
	return &GrpcServer{
		googleGrpcServer: googleGrpcServer,
		protoServer:      server,
	}, nil
}
