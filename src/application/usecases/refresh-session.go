package usecases

import (
	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/application/providers"
	"github.com/AndreyArthur/oganessone/src/application/repositories"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type RefreshSessionUseCase struct {
	repository repositories.AccountsRepository
	cache      providers.CacheProvider
	session    providers.SessionProvider
}

func (refreshSessionUseCase *RefreshSessionUseCase) Execute(
	data *definitions.RefreshSessionDTO,
) (*definitions.RefreshSessionResult, *shared.Error) {
	accountId, _ := refreshSessionUseCase.cache.Get(data.SessionKey)
	account, _ := refreshSessionUseCase.repository.FindById(accountId)
	sessionData, _ := refreshSessionUseCase.session.Generate(account.Id)
	return &definitions.RefreshSessionResult{
		Account:    account,
		SessionKey: sessionData.Key,
	}, nil
}

func NewRefreshSessionUseCase(
	repository repositories.AccountsRepository,
	cache providers.CacheProvider,
	session providers.SessionProvider,
) (*RefreshSessionUseCase, *shared.Error) {
	return &RefreshSessionUseCase{
		repository: repository,
		cache:      cache,
		session:    session,
	}, nil
}
