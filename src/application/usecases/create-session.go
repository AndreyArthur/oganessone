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
	repository repositories.UsersRepository
	encrypter  providers.EncrypterProvider
	session    providers.SessionProvider
}

func (createSessionUseCase *CreateSessionUseCase) findUser(
	login string,
) (*entities.UserEntity, *entities.UserEntity, *shared.Error) {
	foundByUsernameChannel, findByUsernameErrorChannel := make(chan *entities.UserEntity), make(chan *shared.Error)
	foundByEmailChannel, findByEmailErrorChannel := make(chan *entities.UserEntity), make(chan *shared.Error)
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
	foundByUsername, foundByEmail, err := createSessionUseCase.findUser(data.Login)
	if err != nil {
		return nil, err
	}
	if foundByUsername == nil && foundByEmail == nil {
		return nil, exceptions.NewUserLoginFailed()
	}
	var user *entities.UserEntity
	if foundByEmail != nil {
		user = foundByEmail
	}
	if foundByUsername != nil {
		user = foundByUsername
	}
	passwordMatches, err := createSessionUseCase.encrypter.
		Compare(data.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !passwordMatches {
		return nil, exceptions.NewUserLoginFailed()
	}
	sessionKey, err := createSessionUseCase.session.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &definitions.CreateSessionResult{
		User:       user,
		SessionKey: sessionKey,
	}, nil
}

func NewCreateSessionUseCase(
	repository repositories.UsersRepository,
	encrypter providers.EncrypterProvider,
	session providers.SessionProvider,
) (*CreateSessionUseCase, *shared.Error) {
	return &CreateSessionUseCase{
		repository: repository,
		encrypter:  encrypter,
		session:    session,
	}, nil
}
