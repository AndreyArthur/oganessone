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

type RefreshSessionUseCaseTest struct{}

func (*RefreshSessionUseCaseTest) setup(t *testing.T) (*usecases.RefreshSessionUseCase, *mock_repositories.MockAccountsRepository, *mock_providers.MockCacheProvider, *mock_providers.MockSessionProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockAccountsRepository(ctrl)
	cache := mock_providers.NewMockCacheProvider(ctrl)
	session := mock_providers.NewMockSessionProvider(ctrl)
	refreshSessionUseCase, _ := usecases.NewRefreshSessionUseCase(repo, cache, session)
	return refreshSessionUseCase, repo, cache, session, ctrl
}

func TestRefreshSessionUseCase_SuccessCase(t *testing.T) {
	// arrange
	useCase, repo, cache, session, ctrl := (&RefreshSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey := "kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w"
	id, username, email, fakeBcryptHash := "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		"username",
		"account@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	ONE_DAY_IN_SECONDS := int((time.Hour.Milliseconds() * 24) / 1000)
	repoAccount := &entities.AccountEntity{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	cache.EXPECT().
		Get(sessionKey).
		Return(id, nil)
	repo.EXPECT().
		FindById(id).
		Return(repoAccount, nil)
	session.EXPECT().
		Generate(id).
		Return(&providers.SessionData{
			AccountId:               id,
			Key:                     "c^fPj?//P$AkZifWqL$W$D*D$jwq9D$!",
			ExpirationTimeInSeconds: ONE_DAY_IN_SECONDS,
		}, nil)
	// act
	result, err := useCase.Execute(&definitions.RefreshSessionDTO{
		SessionKey: sessionKey,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, result.Account.IsValid())
	assert.NotEqual(t, result.SessionKey, sessionKey)
}

func TestRefreshSessionUseCase_SessionNotFound(t *testing.T) {
	// arrange
	useCase, _, cache, _, ctrl := (&RefreshSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey := "kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w"
	cache.EXPECT().
		Get(sessionKey).
		Return("", nil)
	// act
	result, err := useCase.Execute(&definitions.RefreshSessionDTO{
		SessionKey: sessionKey,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, exceptions.NewSessionNotFound())
}

func TestRefreshSessionUseCase_CacheGetReturnError(t *testing.T) {
	// arrange
	useCase, _, cache, _, ctrl := (&RefreshSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey := "kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w"
	cache.EXPECT().
		Get(sessionKey).
		Return("", &shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.RefreshSessionDTO{
		SessionKey: sessionKey,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}

func TestRefreshSessionUseCase_AccountNotFound(t *testing.T) {
	// arrange
	useCase, repo, cache, _, ctrl := (&RefreshSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey, accountId :=
		"kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w",
		"9b157773-fbb4-d04c-9de6-d086cf37d7c7"
	cache.EXPECT().
		Get(sessionKey).
		Return(accountId, nil)
	repo.EXPECT().
		FindById(accountId).
		Return(nil, nil)
	// act
	result, err := useCase.Execute(&definitions.RefreshSessionDTO{
		SessionKey: sessionKey,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, exceptions.NewAccountNotFound())
}

func TestRefreshSessionUseCase_RepositoryFindByIdReturnError(t *testing.T) {
	// arrange
	useCase, repo, cache, _, ctrl := (&RefreshSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey, accountId :=
		"kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w",
		"9b157773-fbb4-d04c-9de6-d086cf37d7c7"
	cache.EXPECT().
		Get(sessionKey).
		Return(accountId, nil)
	repo.EXPECT().
		FindById(accountId).
		Return(nil, &shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.RefreshSessionDTO{
		SessionKey: sessionKey,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}

func TestRefreshSessionUseCase_SessionGenerateReturnError(t *testing.T) {
	// arrange
	useCase, repo, cache, session, ctrl := (&RefreshSessionUseCaseTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey := "kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w"
	id, username, email, fakeBcryptHash := "9b157773-fbb4-d04c-9de6-d086cf37d7c7",
		"username",
		"account@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	repoAccount := &entities.AccountEntity{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  fakeBcryptHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	cache.EXPECT().
		Get(sessionKey).
		Return(id, nil)
	repo.EXPECT().
		FindById(id).
		Return(repoAccount, nil)
	session.EXPECT().
		Generate(id).
		Return(nil, &shared.Error{})
	// act
	result, err := useCase.Execute(&definitions.RefreshSessionDTO{
		SessionKey: sessionKey,
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}
