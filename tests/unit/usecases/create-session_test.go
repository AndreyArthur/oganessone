package test_usecases

import (
	"testing"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/oganessone/src/application/usecases"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type CreateSessionUseCaseTest struct{}

func (*CreateSessionUseCaseTest) setup(t *testing.T) (*usecases.CreateSessionUseCase, *mock_repositories.MockUsersRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createSessionUseCase, _ := usecases.NewCreateSessionUseCase(repo)
	return createSessionUseCase, repo, ctrl
}

func TestCreateSessionUseCase_UserNotFound(t *testing.T) {
	// arrange
	useCase, repo, ctrl := (&CreateSessionUseCaseTest{}).setup(t)
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
