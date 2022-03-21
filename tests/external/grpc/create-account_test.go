package test_grpc

import (
	"context"
	"database/sql"
	"log"
	"net"
	"testing"

	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/tests/helpers/verifier"
	"github.com/stretchr/testify/assert"
	google_grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateAccountGrpcTest struct{}

func (*CreateAccountGrpcTest) setup() (protobuf.AccountsServiceClient, func(), *sql.DB) {
	env, err := helpers.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = env.Load("test")
	if err != nil {
		log.Fatal(err)
	}
	db, _ := database.NewDatabase()
	sql, _ := db.Connect()
	lis, goerr := net.Listen("tcp", "0.0.0.0:50051")
	if goerr != nil {
		log.Fatal(goerr)
	}
	googleGrpcServer := google_grpc.NewServer()
	server, err := grpc.NewGrpcServer(googleGrpcServer)
	if err != nil {
		log.Fatal(err)
	}
	startedChannel := make(chan bool)
	go func() {
		startedChannel <- true
		server.Start(lis)
	}()
	started := <-startedChannel
	if !started {
		log.Fatal("failed to start the server")
	}
	connection, goerr := google_grpc.Dial("localhost:50051", google_grpc.WithTransportCredentials(insecure.NewCredentials()))
	if goerr != nil {
		log.Fatal(goerr)
	}
	client := protobuf.NewAccountsServiceClient(connection)
	closeConnections := func() {
		connection.Close()
		googleGrpcServer.Stop()
	}
	return client, closeConnections, sql
}

func TestGrpcCreateAccount_Success(t *testing.T) {
	// arrange
	client, closeConnections, sql := (&CreateAccountGrpcTest{}).setup()
	defer closeConnections()
	defer sql.Query("DELETE FROM accounts;")
	username, email, password := "username", "account@email.com", "p4ssword"
	// act
	response, goerr := client.CreateAccount(context.Background(), &protobuf.CreateAccountRequest{
		Data: &protobuf.CreateAccountData{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Error)
	assert.True(t, verifier.IsUuid(response.Data.Account.Id))
	assert.True(t, verifier.IsAccountUsername(response.Data.Account.Username))
	assert.True(t, verifier.IsEmail(response.Data.Account.Email))
	assert.True(t, verifier.IsISO8601(response.Data.Account.CreatedAt))
	assert.True(t, verifier.IsISO8601(response.Data.Account.UpdatedAt))
}

func TestGrpcCreateAccount_UsernameAlreadyInUse(t *testing.T) {
	// arrange
	client, closeConnections, sql := (&CreateAccountGrpcTest{}).setup()
	defer closeConnections()
	defer sql.Query("DELETE FROM accounts;")
	username, email, password := "username", "account@email.com", "p4ssword"
	sql.Query(`
		INSERT INTO accounts (
			username, email, password
		) VALUES ( $1, $2, $3 );
	`, username, "other@email.com", "$2y$10$hRAVNUr.t6UpY1J0bQKmhO5x/K9rZPOGAPdx3HICkCrOUHR/3eyxW")
	// act
	response, goerr := client.CreateAccount(context.Background(), &protobuf.CreateAccountRequest{
		Data: &protobuf.CreateAccountData{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Data)
	assert.Equal(t, response.Error.Name, "AccountUsernameAlreadyInUse")
}

func TestGrpcCreateAccount_EmailAlreadyInUse(t *testing.T) {
	// arrange
	client, closeConnections, sql := (&CreateAccountGrpcTest{}).setup()
	defer closeConnections()
	defer sql.Query("DELETE FROM accounts;")
	username, email, password := "username", "account@email.com", "p4ssword"
	sql.Query(`
		INSERT INTO accounts (
			username, email, password
		) VALUES ( $1, $2, $3 );
	`, "other_username", email, "$2y$10$hRAVNUr.t6UpY1J0bQKmhO5x/K9rZPOGAPdx3HICkCrOUHR/3eyxW")
	// act
	response, goerr := client.CreateAccount(context.Background(), &protobuf.CreateAccountRequest{
		Data: &protobuf.CreateAccountData{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Data)
	assert.Equal(t, response.Error.Name, "AccountEmailAlreadyInUse")
}
