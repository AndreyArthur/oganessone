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

type RefreshSessionGrpcTest struct{}

func (*RefreshSessionGrpcTest) setup() (protobuf.AccountsServiceClient, func(), *sql.DB, redis.Conn) {
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

func TestRefreshSession_SuccessCase(t *testing.T) {
	// arrange
	client, closeConnections, sql, redis := (&RefreshSessionGrpcTest{}).setup()
	defer closeConnections()
	defer redis.Do("FLUSHALL")
	defer sql.Query("DELETE FROM accounts;")
	encrypter, _ := adapters.NewEncrypterAdapter()
	password := "p4ssword"
	uuid, _ := helpers.NewUuid()
	session, _ := adapters.NewSessionAdapter()
	cache, _ := adapters.NewCacheAdapter(redis)
	hash, _ := encrypter.Hash(password)
	id, username, email :=
		uuid.Generate(),
		"username",
		"account@email.com"
	sql.Query(`
		INSERT INTO accounts (
			id, username, email, password
		) VALUES ( $1, $2, $3, $4 );
	`, id, username, email, hash)
	accountSession, _ := session.Generate(id)
	cache.Set(
		accountSession.Key,
		accountSession.AccountId,
		accountSession.ExpirationTimeInSeconds,
	)
	// act
	response, goerr := client.RefreshSession(context.Background(), &protobuf.RefreshSessionRequest{
		Data: &protobuf.RefreshSessionRequestData{
			SessionKey: accountSession.Key,
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

func TestRefreshSession_SessionNotFound(t *testing.T) {
	// arrange
	client, closeConnections, sql, redis := (&RefreshSessionGrpcTest{}).setup()
	defer closeConnections()
	defer redis.Do("FLUSHALL")
	defer sql.Query("DELETE FROM accounts;")
	encrypter, _ := adapters.NewEncrypterAdapter()
	password := "p4ssword"
	uuid, _ := helpers.NewUuid()
	session, _ := adapters.NewSessionAdapter()
	hash, _ := encrypter.Hash(password)
	id, username, email :=
		uuid.Generate(),
		"username",
		"account@email.com"
	sql.Query(`
		INSERT INTO accounts (
			id, username, email, password
		) VALUES ( $1, $2, $3, $4 );
	`, id, username, email, hash)
	accountSession, _ := session.Generate(id)
	// act
	response, goerr := client.RefreshSession(context.Background(), &protobuf.RefreshSessionRequest{
		Data: &protobuf.RefreshSessionRequestData{
			SessionKey: accountSession.Key,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Data)
	assert.Equal(t, response.Error.Message, exceptions.NewSessionNotFound().Message)
}

func TestRefreshSession_AccountNotFound(t *testing.T) {
	// arrange
	client, closeConnections, sql, redis := (&RefreshSessionGrpcTest{}).setup()
	defer closeConnections()
	defer redis.Do("FLUSHALL")
	defer sql.Query("DELETE FROM accounts;")
	uuid, _ := helpers.NewUuid()
	session, _ := adapters.NewSessionAdapter()
	cache, _ := adapters.NewCacheAdapter(redis)
	id := uuid.Generate()
	accountSession, _ := session.Generate(id)
	cache.Set(
		accountSession.Key,
		accountSession.AccountId,
		accountSession.ExpirationTimeInSeconds,
	)
	// act
	response, goerr := client.RefreshSession(context.Background(), &protobuf.RefreshSessionRequest{
		Data: &protobuf.RefreshSessionRequestData{
			SessionKey: accountSession.Key,
		},
	})
	// assert
	assert.Nil(t, goerr)
	assert.Nil(t, response.Data)
	assert.Equal(t, response.Error.Message, exceptions.NewAccountNotFound().Message)
}
