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

type CreateAccountUseCase struct {
	repository repositories.AccountsRepository
	encrypter  providers.EncrypterProvider
}

func (createAccountUseCase *CreateAccountUseCase) sanitize(
	username *string, email *string, password *string,
) {
	*username = strings.TrimSpace(*username)
	*email = strings.TrimSpace(*email)
	*password = strings.TrimSpace(*password)
}

func (createAccountUseCase *CreateAccountUseCase) findAccount(
	username string, email string,
) (*entities.AccountEntity, *entities.AccountEntity, *shared.Error) {
	foundByUsernameChannel, findByUsernameErrorChannel := make(chan *entities.AccountEntity), make(chan *shared.Error)
	foundByEmailChannel, findByEmailErrorChannel := make(chan *entities.AccountEntity), make(chan *shared.Error)
	go func() {
		foundByUsername, err := createAccountUseCase.repository.FindByUsername(
			username, false,
		)
		foundByUsernameChannel <- foundByUsername
		findByUsernameErrorChannel <- err
	}()
	go func() {
		foundByEmail, err := createAccountUseCase.repository.FindByEmail(email)
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

func (createAccountUseCase *CreateAccountUseCase) Execute(
	data *definitions.CreateAccountDTO,
) (*definitions.CreateAccountResult, *shared.Error) {
	createAccountUseCase.sanitize(&data.Username, &data.Email, &data.Password)
	foundByUsername, foundByEmail, err := createAccountUseCase.
		findAccount(data.Username, data.Email)
	if err != nil {
		return nil, err
	}
	if foundByUsername != nil {
		return nil, exceptions.NewAccountUsernameAlreadyInUse()
	}
	if foundByEmail != nil {
		return nil, exceptions.NewAccountEmailAlreadyInUse()
	}
	hashedPassword, err := createAccountUseCase.encrypter.Hash(data.Password)
	if err != nil {
		return nil, exceptions.NewInternalServerError()
	}
	account, err := createAccountUseCase.repository.Create(&dtos.AccountDTO{
		Username: data.Username,
		Email:    data.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	err = account.IsPasswordValid(data.Password)
	if err != nil {
		return nil, err
	}
	err = createAccountUseCase.repository.Save(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func NewCreateAccountUseCase(
	repository repositories.AccountsRepository,
	encrypter providers.EncrypterProvider,
) (*CreateAccountUseCase, *shared.Error) {
	createAccountUseCase := &CreateAccountUseCase{
		repository: repository,
		encrypter:  encrypter,
	}
	return createAccountUseCase, nil
}
