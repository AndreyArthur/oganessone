package test_grpc

import (
	"context"
	"database/sql"
	"log"
	"net"
	"testing"

	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/infrastructure/adapters"
	"github.com/AndreyArthur/oganessone/src/infrastructure/cache"
	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/tests/helpers/verifier"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	google_grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateSessionGrpcTest struct{}

func (*CreateSessionGrpcTest) setup() (protobuf.AccountsServiceClient, func(), *sql.DB, redis.Conn) {
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
	redis, _ := cache.NewRedis()
	redisConnection, _ := redis.Connect()
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
		redisConnection.Close()
		googleGrpcServer.Stop()
	}
	return client, closeConnections, sql, redisConnection
}

func TestCreateSession_SuccessByUsername(t *testing.T) {
	// arrange
	client, closeConnections, sql, redis := (&CreateSessionGrpcTest{}).setup()
	defer closeConnections()
	defer redis.Do("FLUSHALL")
	defer sql.Query("DELETE FROM accounts;")
	encrypter, _ := adapters.NewEncrypterAdapter()
	password := "p4ssword"
	hash, _ := encrypter.Hash(password)
	username, email :=
		"username",
		"account@email.com"
	sql.Query(`
		INSERT INTO accounts (
			username, email, password
		) VALUES ( $1, $2, $3 );
	`, username, email, hash)
	// act
	response, goerr := client.CreateSession(context.Background(), &protobuf.CreateSessionRequest{
		Data: &protobuf.CreateSessionRequestData{
			Login:    username,
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

func TestCreateSession_SuccessByEmail(t *testing.T) {
	// arrange
	client, closeConnections, sql, redis := (&CreateSessionGrpcTest{}).setup()
	defer closeConnections()
	defer redis.Do("FLUSHALL")
	defer sql.Query("DELETE FROM accounts;")
	encrypter, _ := adapters.NewEncrypterAdapter()
	password := "p4ssword"
	hash, _ := encrypter.Hash(password)
	username, email :=
		"username",
		"account@email.com"
	sql.Query(`
		INSERT INTO accounts (
			username, email, password
		) VALUES ( $1, $2, $3 );
	`, username, email, hash)
	// act
	response, goerr := client.CreateSession(context.Background(), &protobuf.CreateSessionRequest{
		Data: &protobuf.CreateSessionRequestData{
			Login:    email,
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

func TestCreateSession_NotFound(t *testing.T) {
	// arrange
	client, closeConnections, _, _ := (&CreateSessionGrpcTest{}).setup()
	defer closeConnections()
	username, password := "username", "p4ssword"
	response, goerr := client.CreateSession(context.Background(), &protobuf.CreateSessionRequest{
		Data: &protobuf.CreateSessionRequestData{
			Login:    username,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Data)
	assert.Equal(t, response.Error.Message, exceptions.NewAccountLoginFailed().Message)
}

func TestCreateSession_PasswordDoesNotMatch(t *testing.T) {
	// arrange
	client, closeConnections, sql, redis := (&CreateSessionGrpcTest{}).setup()
	defer closeConnections()
	defer redis.Do("FLUSHALL")
	defer sql.Query("DELETE FROM accounts;")
	encrypter, _ := adapters.NewEncrypterAdapter()
	hash, _ := encrypter.Hash("0ther password")
	username, email, password :=
		"username",
		"account@email.com",
		"p4ssword"
	sql.Query(`
		INSERT INTO accounts (
			username, email, password
		) VALUES ( $1, $2, $3 );
	`, username, email, hash)
	// act
	response, goerr := client.CreateSession(context.Background(), &protobuf.CreateSessionRequest{
		Data: &protobuf.CreateSessionRequestData{
			Login:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Data)
	assert.Equal(t, response.Error.Message, exceptions.NewAccountLoginFailed().Message)
}
