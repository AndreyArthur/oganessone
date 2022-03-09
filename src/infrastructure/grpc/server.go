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

func (*server) error(err *shared.Error) *protobuf.Error {
	return &protobuf.Error{
		Type:    err.Type,
		Name:    err.Name,
		Message: err.Message,
	}
}

func (srvr *server) CreateAccount(
	ctx context.Context, request *protobuf.CreateAccountRequest,
) (*protobuf.CreateAccountResponse, error) {
	rpc, err := factories.MakeCreateAccountRpc()
	if err != nil {
		return &protobuf.CreateAccountResponse{
			Error: srvr.error(err),
			Data:  nil,
		}, nil
	}
	return rpc.Perform(ctx, request)
}

func (srvr *server) CreateSession(
	ctx context.Context, request *protobuf.CreateSessionRequest,
) (*protobuf.CreateSessionResponse, error) {
	rpc, err := factories.MakeCreateSessionRpc()
	if err != nil {
		return &protobuf.CreateSessionResponse{
			Error: srvr.error(err),
			Data:  nil,
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
