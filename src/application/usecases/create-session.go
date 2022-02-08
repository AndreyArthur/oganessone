package usecases

import (
	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/application/repositories"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateSessionUseCase struct {
	repository repositories.UsersRepository
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
	return nil, nil
}

func NewCreateSessionUseCase(repository repositories.UsersRepository) (*CreateSessionUseCase, *shared.Error) {
	return &CreateSessionUseCase{repository: repository}, nil
}
