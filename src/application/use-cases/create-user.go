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
	foundByUsername, _ := createUserUseCase.repository.FindByUsername(
		username, true,
	)
	if foundByUsername != nil {
		return nil, exceptions.NewUserUsernameAlreadyInUse()
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
