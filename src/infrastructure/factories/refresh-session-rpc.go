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

func MakeRefreshSessionRpc() (*rpcs.RefreshSessionRpc, *shared.Error) {
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
	refreshSessionUseCase, err := usecases.NewRefreshSessionUseCase(
		repo,
		cache,
		session,
	)
	if err != nil {
		return nil, err
	}
	refreshSessionPresenter, err := presenters.NewRefreshSessionPresenter(
		refreshSessionUseCase,
	)
	if err != nil {
		return nil, err
	}
	refreshSessionRpc, err := rpcs.NewRefreshSessionRpc(refreshSessionPresenter)
	if err != nil {
		return nil, err
	}
	return refreshSessionRpc, nil
}
