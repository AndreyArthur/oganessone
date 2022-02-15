package factories

import (
	usecases "github.com/AndreyArthur/oganessone/src/application/usecases"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/adapters"
	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/rpcs"
	"github.com/AndreyArthur/oganessone/src/infrastructure/repositories"
	"github.com/AndreyArthur/oganessone/src/presentation/presenters"
)

func MakeCreateUserRpc() (*rpcs.CreateUserRpc, *shared.Error) {
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
	createUser, err := usecases.NewCreateAccountUseCase(repo, encrypter)
	if err != nil {
		return nil, err
	}
	createUserPresenter, err := presenters.NewCreateAccountPresenter(createUser)
	if err != nil {
		return nil, err
	}
	createUserRpc, err := rpcs.NewCreateUserRpc(createUserPresenter)
	if err != nil {
		return nil, err
	}
	return createUserRpc, nil
}
