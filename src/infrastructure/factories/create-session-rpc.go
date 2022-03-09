package factories

import (
	"github.com/AndreyArthur/oganessone/src/application/usecases"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/adapters"
	"github.com/AndreyArthur/oganessone/src/infrastructure/cache"
	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/rpcs"
	"github.com/AndreyArthur/oganessone/src/infrastructure/repositories"
	"github.com/AndreyArthur/oganessone/src/presentation/presenters"
)

func MakeCreateSessionRpc() (*rpcs.CreateSessionRpc, *shared.Error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}
	sql, err := db.Connect()
	if err != nil {
		return nil, err
	}
	repo, err := repositories.NewAccountsRepositoryPostgres(sql)
	if err != nil {
		return nil, err
	}
	encrypter, err := adapters.NewEncrypterAdapter()
	if err != nil {
		return nil, err
	}
	session, err := adapters.NewSessionAdapter()
	if err != nil {
		return nil, err
	}
	redis, err := cache.NewRedis()
	if err != nil {
		return nil, err
	}
	connection, err := redis.Connect()
	if err != nil {
		return nil, err
	}
	cache, err := adapters.NewCacheAdapter(connection)
	if err != nil {
		return nil, err
	}
	createSessionUseCase, err := usecases.NewCreateSessionUseCase(
		repo,
		encrypter,
		session,
		cache,
	)
	if err != nil {
		return nil, err
	}
	createSessionPresenter, err := presenters.NewCreateSessionPresenter(
		createSessionUseCase,
	)
	if err != nil {
		return nil, err
	}
	createSessionRpc, err := rpcs.NewCreateSessionRpc(createSessionPresenter)
	if err != nil {
		return nil, err
	}
	return createSessionRpc, nil
}
