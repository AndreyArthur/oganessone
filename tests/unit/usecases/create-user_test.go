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

type CreateUserUseCaseTest struct{}

func (*CreateUserUseCaseTest) setup(t *testing.T) (*usecases.CreateAccountUseCase, *mock_repositories.MockAccountsRepository, *mock_providers.MockEncrypterProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockAccountsRepository(ctrl)
	encrypter := mock_providers.NewMockEncrypterProvider(ctrl)
	createUserUseCase, _ := usecases.NewCreateAccountUseCase(repo, encrypter)
	return createUserUseCase, repo, encrypter, ctrl
}

func TestCreateUserUseCase_SuccessCase(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoUser := &entities.AccountEntity{
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
		Return(repoUser, nil)
	repo.EXPECT().
		Save(repoUser).
		Return(nil)
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user.IsValid())
}

func TestCreateUserUseCase_SanitizeValues(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "  username ", "  user@email.com ", " p4ssword  "
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoUser := &entities.AccountEntity{
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
		Return(repoUser, nil)
	repo.EXPECT().
		Save(repoUser).
		Return(nil)
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.NotEqual(t, user.Username, username)
	assert.Equal(t, user.Username, strings.TrimSpace(username))
	assert.NotEqual(t, user.Email, email)
	assert.Equal(t, user.Email, strings.TrimSpace(email))
}

func TestCreateUserUseCase_FoundByUsername(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(&entities.AccountEntity{}, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, exceptions.NewAccountUsernameAlreadyInUse())
	assert.Nil(t, user)
}

func TestCreateUserUseCase_FindByUsernameReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, &shared.Error{})
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.AccountEntity{}, nil)
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, user)
}

func TestCreateUserUseCase_FoundByEmail(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.AccountEntity{}, nil)
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, user)
	assert.Equal(t, err, exceptions.NewAccountEmailAlreadyInUse())
}

func TestCreateUserUseCase_FindByEmailReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, &shared.Error{})
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, user)
}

func TestCreateUserUseCase_EncrypterHashReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
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
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, exceptions.NewInternalServerError())
	assert.Nil(t, user)
}

func TestCreateUserUseCase_CreateReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
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
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, user)
}

func TestCreateUserUseCase_PasswordValidationReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, invalidPassword := "username", "user@email.com", "invalid password"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoUser := &entities.AccountEntity{
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
		Return(repoUser, nil)
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: invalidPassword,
	})
	// assert
	assert.Nil(t, user)
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())
}

func TestCreateUserUseCase_SaveReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := (&CreateUserUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoUser := &entities.AccountEntity{
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
		Return(repoUser, nil)
	repo.EXPECT().
		Save(repoUser).
		Return(&shared.Error{})
	// act
	user, err := useCase.Execute(&definitions.CreateAccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, user)
	assert.Equal(t, err, &shared.Error{})
}
