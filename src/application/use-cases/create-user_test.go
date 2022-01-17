package usecases

import (
	"errors"
	"testing"

	mock_providers "github.com/AndreyArthur/oganessone/src/application/providers/mocks"
	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*CreateUserUseCase, *mock_repositories.MockUsersRepository, *mock_providers.MockEncrypterProvider, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	encrypter := mock_providers.NewMockEncrypterProvider(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo, encrypter)

	return createUserUseCase, repo, encrypter, ctrl
}

func TestCreateUserUseCase_NotFoundByUsername(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return("$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW", nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Email, email)
}

func TestCreateUserUseCase_FoundByUsername(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(&entities.UserEntity{}, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, exceptions.NewUserUsernameAlreadyInUse())
	assert.Nil(t, user)
}

func TestCreateUserUseCase_FindByUsernameReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, errors.New(""))
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.UserEntity{}, nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, exceptions.NewInternalServerError())
	assert.Nil(t, user)
}

func TestCreateUserUseCase_NotFoundByEmail(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return("$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW", nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Email, email)
}

func TestCreateUserUseCase_FoundByEmail(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.UserEntity{}, nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Nil(t, user)
	assert.Equal(t, err, exceptions.NewUserEmailAlreadyInUse())
}

func TestCreateUserUseCase_FindByEmailReturnError(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, errors.New(""))
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, exceptions.NewInternalServerError())
	assert.Nil(t, user)
}
