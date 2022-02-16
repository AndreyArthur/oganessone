package test_usecases

import (
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/application/providers"
	mock_providers "github.com/AndreyArthur/oganessone/src/application/providers/mocks"
	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/oganessone/src/application/usecases"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type CreateSessionUseCaseTest struct{}

func (*CreateSessionUseCaseTest) setup(t *testing.T) (*usecases.CreateSessionUseCase, *mock_repositories.MockAccountsRepository, *mock_providers.MockEncrypterProvider, *mock_providers.MockSessionProvider, *mock_providers.MockCacheProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockAccountsRepository(ctrl)
	encrypter := mock_providers.NewMockEncrypterProvider(ctrl)
	session := mock_providers.NewMockSessionProvider(ctrl)
	cache := mock_providers.NewMockCacheProvider(ctrl)
	createSessionUseCase, _ := usecases.NewCreateSessionUseCase(repo, encrypter, session, cache)
	return createSessionUseCase, repo, encrypter, session, cache, ctrl
}

func TestCreateSessionUseCase_SuccessCaseByUsername(t *testing.T) {
	// arrange
	useCase, repo, encrypter, session, cache, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
	sessionKey := "session_key_example"
	ONE_DAY_IN_SECONDS := int((time.Hour.Milliseconds() * 24) / 1000)
	repo.EXPECT().
		FindByEmail(username).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(username, true).
		Return(repoAccount, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(true, nil)
	session.EXPECT().
		Generate(repoAccount.Id).
		Return(&providers.SessionData{
			Key:                     sessionKey,
			AccountId:               repoAccount.Id,
			ExpirationTimeInSeconds: ONE_DAY_IN_SECONDS,
		}, nil)
	cache.EXPECT().
		Set(sessionKey, repoAccount.Id, ONE_DAY_IN_SECONDS).
		Return(nil)
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    username,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, result.Account.IsValid())
	assert.Equal(t, result.SessionKey, sessionKey)
}

func TestCreateSessionUseCase_SuccessCaseByEmail(t *testing.T) {
	// arrange
	useCase, repo, encrypter, session, cache, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
	sessionKey := "session_key_example"
	ONE_DAY_IN_SECONDS := int((time.Hour.Milliseconds() * 24) / 1000)
	repo.EXPECT().
		FindByEmail(email).
		Return(repoAccount, nil)
	repo.EXPECT().
		FindByUsername(email, true).
		Return(nil, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(true, nil)
	session.EXPECT().
		Generate(repoAccount.Id).
		Return(&providers.SessionData{
			Key:                     sessionKey,
			AccountId:               repoAccount.Id,
			ExpirationTimeInSeconds: ONE_DAY_IN_SECONDS,
		}, nil)
	cache.EXPECT().
		Set(sessionKey, repoAccount.Id, ONE_DAY_IN_SECONDS).
		Return(nil)
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, result.Account.IsValid())
	assert.Equal(t, result.SessionKey, sessionKey)
}

func TestCreateSessionUseCase_FindByUsernameReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, _, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	login, password := "username", "p4ssword"
	repo.EXPECT().
		FindByEmail(login).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(login, true).
		Return(nil, &shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    login,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}

func TestCreateSessionUseCase_FindByEmailReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, _, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	login, password := "username", "p4ssword"
	repo.EXPECT().
		FindByEmail(login).
		Return(nil, &shared.Error{})
	repo.EXPECT().
		FindByUsername(login, true).
		Return(nil, nil)
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    login,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}

func TestCreateSessionUseCase_AccountNotFound(t *testing.T) {
	// arrange
	useCase, repo, _, _, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	login, password := "username", "p4ssword"
	repo.EXPECT().
		FindByEmail(login).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(login, true).
		Return(nil, nil)
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    login,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, exceptions.NewAccountLoginFailed())
}

func TestCreateSessionUseCase_EncrypterCompareReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, _, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
		FindByEmail(username).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(username, true).
		Return(repoAccount, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(true, &shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    username,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}

func TestCreateSessionUseCase_AccountPasswordDoesNotMatch(t *testing.T) {
	// arrange
	useCase, repo, encrypter, _, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
		FindByEmail(username).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(username, true).
		Return(repoAccount, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(false, nil)
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    username,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, exceptions.NewAccountLoginFailed())
}

func TestCreateSessionUseCase_SessionGenerateKeyReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, session, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
	repo.EXPECT().FindByEmail(username).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(username, true).
		Return(repoAccount, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(true, nil)
	session.EXPECT().
		Generate(repoAccount.Id).
		Return(nil, &shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    username,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}

func TestCreateSessionUseCase_FirstCacheSetReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, session, cache, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
	sessionKey := "session_key_example"
	ONE_DAY_IN_SECONDS := int((time.Hour.Milliseconds() * 24) / 1000)
	repo.EXPECT().
		FindByEmail(username).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(username, true).
		Return(repoAccount, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(true, nil)
	session.EXPECT().
		Generate(repoAccount.Id).
		Return(&providers.SessionData{
			Key:                     sessionKey,
			AccountId:               repoAccount.Id,
			ExpirationTimeInSeconds: ONE_DAY_IN_SECONDS,
		}, nil)
	cache.EXPECT().
		Set(sessionKey, repoAccount.Id, ONE_DAY_IN_SECONDS).
		Return(&shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    username,
		Password: password,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}
