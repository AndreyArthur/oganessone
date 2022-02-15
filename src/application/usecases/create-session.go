package usecases

import (
	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/application/providers"
	"github.com/AndreyArthur/oganessone/src/application/repositories"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateSessionUseCase struct {
	repository repositories.AccountsRepository
	encrypter  providers.EncrypterProvider
	session    providers.SessionProvider
	cache      providers.CacheProvider
}

func (createSessionUseCase *CreateSessionUseCase) findAccount(
	login string,
) (*entities.AccountEntity, *entities.AccountEntity, *shared.Error) {
	foundByUsernameChannel, findByUsernameErrorChannel := make(chan *entities.AccountEntity), make(chan *shared.Error)
	foundByEmailChannel, findByEmailErrorChannel := make(chan *entities.AccountEntity), make(chan *shared.Error)
	go func() {
		foundByUsername, err := createSessionUseCase.repository.FindByUsername(
			login, true,
		)
		foundByUsernameChannel <- foundByUsername
		findByUsernameErrorChannel <- err
	}()
	go func() {
		foundByEmail, err := createSessionUseCase.repository.FindByEmail(login)
		foundByEmailChannel <- foundByEmail
		findByEmailErrorChannel <- err
	}()
	foundByUsername, foundByEmail := <-foundByUsernameChannel, <-foundByEmailChannel
	findByUsernameError, findByEmailError := <-findByUsernameErrorChannel, <-findByEmailErrorChannel
	if findByUsernameError != nil {
		return nil, nil, findByUsernameError
	}
	if findByEmailError != nil {
		return nil, nil, findByEmailError
	}
	return foundByUsername, foundByEmail, nil
}

func (createSessionUseCase *CreateSessionUseCase) Execute(
	data *definitions.CreateSessionDTO,
) (*definitions.CreateSessionResult, *shared.Error) {
	foundByUsername, foundByEmail, err := createSessionUseCase.findAccount(data.Login)
	if err != nil {
		return nil, err
	}
	if foundByUsername == nil && foundByEmail == nil {
		return nil, exceptions.NewAccountLoginFailed()
	}
	var account *entities.AccountEntity
	if foundByEmail != nil {
		account = foundByEmail
	}
	if foundByUsername != nil {
		account = foundByUsername
	}
	passwordMatches, err := createSessionUseCase.encrypter.
		Compare(data.Password, account.Password)
	if err != nil {
		return nil, err
	}
	if !passwordMatches {
		return nil, exceptions.NewAccountLoginFailed()
	}
	sessionData, err := createSessionUseCase.session.Generate(account.Id)
	if err != nil {
		return nil, err
	}
	err = createSessionUseCase.cache.Set(
		sessionData.Key,
		sessionData.AccountId,
		sessionData.ExpirationTimeInSeconds,
	)
	if err != nil {
		return nil, err
	}
	return &definitions.CreateSessionResult{
		Account:    account,
		SessionKey: sessionData.Key,
	}, nil
}

func NewCreateSessionUseCase(
	repository repositories.AccountsRepository,
	encrypter providers.EncrypterProvider,
	session providers.SessionProvider,
	cache providers.CacheProvider,
) (*CreateSessionUseCase, *shared.Error) {
	return &CreateSessionUseCase{
		repository: repository,
		encrypter:  encrypter,
		session:    session,
		cache:      cache,
	}, nil
}
