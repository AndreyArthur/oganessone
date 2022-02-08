package test_usecases

import (
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	mock_providers "github.com/AndreyArthur/oganessone/src/application/providers/mocks"
	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/oganessone/src/application/usecases"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type CreateSessionUseCaseTest struct{}

func (*CreateSessionUseCaseTest) setup(t *testing.T) (*usecases.CreateSessionUseCase, *mock_repositories.MockUsersRepository, *mock_providers.MockEncrypterProvider, *mock_providers.MockSessionProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	encrypter := mock_providers.NewMockEncrypterProvider(ctrl)
	session := mock_providers.NewMockSessionProvider(ctrl)
	createSessionUseCase, _ := usecases.NewCreateSessionUseCase(repo, encrypter, session)
	return createSessionUseCase, repo, encrypter, session, ctrl
}

func TestCreateSessionUseCase_SuccessCase(t *testing.T) {
	// arrange
	useCase, repo, encrypter, session, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoUser := &entities.UserEntity{
		Id:        "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		Username:  username,
		Email:     email,
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sessionKey := "session_key_example"
	repo.EXPECT().
		FindByEmail(username).
		Return(nil, nil)
	repo.EXPECT().
		FindByUsername(username, true).
		Return(repoUser, nil)
	encrypter.EXPECT().
		Compare(password, fakeBcryptHash).
		Return(true, nil)
	session.EXPECT().
		GenerateKey().
		Return(sessionKey, nil)
	// act
	result, err := useCase.Execute(&definitions.CreateSessionDTO{
		Login:    username,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, result.User.IsValid())
	assert.Equal(t, result.SessionKey, sessionKey)
}

func TestCreateSessionUseCase_UserNotFound(t *testing.T) {
	// arrange
	useCase, repo, _, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
	assert.Equal(t, err, exceptions.NewUserLoginFailed())
}

func TestCreateSessionUseCase_UserPasswordDoesNotMatch(t *testing.T) {
	// arrange
	useCase, repo, encrypter, _, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	fakeBcryptHash := "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoUser := &entities.UserEntity{
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
		Return(repoUser, nil)
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
	assert.Equal(t, err, exceptions.NewUserLoginFailed())
}