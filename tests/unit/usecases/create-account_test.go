package test_usecases

import (
	"strings"
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	mock_providers "github.com/AndreyArthur/oganessone/src/application/providers/mocks"
	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	usecases "github.com/AndreyArthur/oganessone/src/application/usecases"
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type CreateAccountUseCaseTest struct{}

func (*CreateAccountUseCaseTest) setup(t *testing.T) (*usecases.CreateAccountUseCase, *mock_repositories.MockAccountsRepository, *mock_providers.MockEncrypterProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockAccountsRepository(ctrl)
	encrypter := mock_providers.NewMockEncrypterProvider(ctrl)
	createAccountUseCase, _ := usecases.NewCreateAccountUseCase(repo, encrypter)
	return createAccountUseCase, repo, encrypter, ctrl
}

func TestCreateAccountUseCase_SuccessCase(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoAccount := &entities.AccountEntity{
		Id:        "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		Username:  username,
		Email:     email,
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(&dtos.AccountDTO{Username: username, Email: email, Password: fakeBcryptHash}).
		Return(repoAccount, nil)
	repo.EXPECT().
		Save(repoAccount).
		Return(nil)
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account.IsValid())
}

func TestCreateAccountUseCase_SanitizeValues(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "  username ", "  account@email.com ", " p4ssword  "
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoAccount := &entities.AccountEntity{
		Id:        "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		Username:  strings.TrimSpace(username),
		Email:     strings.TrimSpace(email),
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.EXPECT().
		FindByUsername(strings.TrimSpace(username), false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(strings.TrimSpace(email)).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(strings.TrimSpace(password)).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(&dtos.AccountDTO{
			Username: strings.TrimSpace(username),
			Email:    strings.TrimSpace(email),
			Password: fakeBcryptHash,
		}).
		Return(repoAccount, nil)
	repo.EXPECT().
		Save(repoAccount).
		Return(nil)
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.NotEqual(t, account.Username, username)
	assert.Equal(t, account.Username, strings.TrimSpace(username))
	assert.NotEqual(t, account.Email, email)
	assert.Equal(t, account.Email, strings.TrimSpace(email))
}

func TestCreateAccountUseCase_FoundByUsername(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(&entities.AccountEntity{}, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, exceptions.NewAccountUsernameAlreadyInUse())
	assert.Nil(t, account)
}

func TestCreateAccountUseCase_FindByUsernameReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, &shared.Error{})
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.AccountEntity{}, nil)
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, account)
}

func TestCreateAccountUseCase_FoundByEmail(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.AccountEntity{}, nil)
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, account)
	assert.Equal(t, err, exceptions.NewAccountEmailAlreadyInUse())
}

func TestCreateAccountUseCase_FindByEmailReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, &shared.Error{})
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, account)
}

func TestCreateAccountUseCase_EncrypterHashReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return("", &shared.Error{})
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, exceptions.NewInternalServerError())
	assert.Nil(t, account)
}

func TestCreateAccountUseCase_CreateReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(&dtos.AccountDTO{Username: username, Email: email, Password: fakeBcryptHash}).
		Return(nil, &shared.Error{})
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, account)
}

func TestCreateAccountUseCase_PasswordValidationReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, invalidPassword := "username", "account@email.com", "invalid password"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoAccount := &entities.AccountEntity{
		Id:        "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		Username:  username,
		Email:     email,
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(invalidPassword).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(&dtos.AccountDTO{Username: username, Email: email, Password: fakeBcryptHash}).
		Return(repoAccount, nil)
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: invalidPassword,
	})
	// assert
	assert.Nil(t, account)
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())
}

func TestCreateAccountUseCase_SaveReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateAccountUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "account@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoAccount := &entities.AccountEntity{
		Id:        "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		Username:  username,
		Email:     email,
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(&dtos.AccountDTO{Username: username, Email: email, Password: fakeBcryptHash}).
		Return(repoAccount, nil)
	repo.EXPECT().
		Save(repoAccount).
		Return(&shared.Error{})
	// act
	account, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, account)
	assert.Equal(t, err, &shared.Error{})
}
