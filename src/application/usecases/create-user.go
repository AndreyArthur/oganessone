package usecases

import (
	"strings"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/application/providers"
	"github.com/AndreyArthur/oganessone/src/application/repositories"
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateUserUseCase struct {
	repository repositories.UsersRepository
	encrypter  providers.EncrypterProvider
}

func (createUserUseCase *CreateUserUseCase) sanitize(
	username *string, email *string, password *string,
) {
	*username = strings.TrimSpace(*username)
	*email = strings.TrimSpace(*email)
	*password = strings.TrimSpace(*password)
}

func (createUserUseCase *CreateUserUseCase) findUser(
	username string, email string,
) (*entities.UserEntity, *entities.UserEntity, *shared.Error) {
	foundByUsernameChannel, findByUsernameErrorChannel := make(chan *entities.UserEntity), make(chan *shared.Error)
	foundByEmailChannel, findByEmailErrorChannel := make(chan *entities.UserEntity), make(chan *shared.Error)
	go func() {
		foundByUsername, err := createUserUseCase.repository.FindByUsername(
			username, false,
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
	if findByUsernameError != nil {
		return nil, nil, findByUsernameError
	}
	if findByEmailError != nil {
		return nil, nil, findByEmailError
	}
	return foundByUsername, foundByEmail, nil
}

func (createUserUseCase *CreateUserUseCase) Execute(
	data *definitions.CreateUserDTO,
) (*definitions.CreateUserResult, *shared.Error) {
	createUserUseCase.sanitize(&data.Username, &data.Email, &data.Password)
	foundByUsername, foundByEmail, err := createUserUseCase.
		findUser(data.Username, data.Email)
	if err != nil {
		return nil, err
	}
	if foundByUsername != nil {
		return nil, exceptions.NewUserUsernameAlreadyInUse()
	}
	if foundByEmail != nil {
		return nil, exceptions.NewUserEmailAlreadyInUse()
	}
	hashedPassword, err := createUserUseCase.encrypter.Hash(data.Password)
	if err != nil {
		return nil, exceptions.NewInternalServerError()
	}
	user, err := createUserUseCase.repository.Create(&dtos.UserDTO{
		Username: data.Username,
		Email:    data.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	err = user.IsPasswordValid(data.Password)
	if err != nil {
		return nil, err
	}
	err = createUserUseCase.repository.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewCreateUserUseCase(
	repository repositories.UsersRepository,
	encrypter providers.EncrypterProvider,
) (*CreateUserUseCase, *shared.Error) {
	createUserUseCase := &CreateUserUseCase{
		repository: repository,
		encrypter:  encrypter,
	}
	return createUserUseCase, nil
}
