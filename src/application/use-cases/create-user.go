package usecases

import (
	"errors"

	"github.com/AndreyArthur/murao-oganessone/src/application/repositories"
	"github.com/AndreyArthur/murao-oganessone/src/core/entities"
)

type CreateUserUseCase struct {
	repository repositories.UsersRepository
}

func (createUserUseCase *CreateUserUseCase) Execute(
	username string, email string, password string,
) (*entities.UserEntity, error) {
	foundByUsername, _ := createUserUseCase.repository.FindByUsername(
		username, true,
	)
	if foundByUsername != nil {
		return nil, errors.New("username is already in use")
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
