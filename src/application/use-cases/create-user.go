package usecases

import (
	"github.com/AndreyArthur/oganessone/src/application/repositories"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateUserUseCase struct {
	repository repositories.UsersRepository
}

func (createUserUseCase *CreateUserUseCase) Execute(
	username string, email string, password string,
) (*entities.UserEntity, *shared.Error) {
	foundByUsernameChannel, findByUsernameErrorChannel := make(chan *entities.UserEntity), make(chan error)
	foundByEmailChannel, findByEmailErrorChannel := make(chan *entities.UserEntity), make(chan error)
	go func() {
		foundByUsername, err := createUserUseCase.repository.FindByUsername(
			username, true,
		)
		foundByUsernameChannel <- foundByUsername
		findByUsernameErrorChannel <- err
	}()
	go func() {
		foundByEmail, err := createUserUseCase.repository.FindByEmail(email)
		foundByEmailChannel <- foundByEmail
		findByEmailErrorChannel <- err
	}()
	foundByUsername, foundByEmail := <-foundByUsernameChannel, <-foundByEmailChannel
	findByUsernameError, findByEmailError := <-findByUsernameErrorChannel, <-findByEmailErrorChannel
	if findByEmailError != nil || findByUsernameError != nil {
		return nil, exceptions.NewInternalServerError()
	}
	if foundByUsername != nil {
		return nil, exceptions.NewUserUsernameAlreadyInUse()
	}
	if foundByEmail != nil {
		return nil, exceptions.NewUserEmailAlreadyInUse()
	}
	return &entities.UserEntity{
		Username: username,
		Email:    email,
	}, nil
}

func NewCreateUserUseCase(
	repository repositories.UsersRepository,
) (*CreateUserUseCase, error) {
	createUserUseCase := &CreateUserUseCase{
		repository: repository,
	}
	return createUserUseCase, nil
}
