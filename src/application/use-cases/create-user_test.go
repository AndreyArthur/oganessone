package usecases

import (
	"testing"

	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_NotFoundByUsername(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo)
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	// act
	user, err := createUserUseCase.Execute(username, email, password)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Email, email)
}

func TestCreateUserUseCase_FoundByUsername(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo)
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(&entities.UserEntity{}, nil)
	// act
	user, err := createUserUseCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, exceptions.NewUserUsernameAlreadyInUse())
	assert.Nil(t, user)
}

func TestCreateUserUseCase_NotFoundByEmail(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo)
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	// act
	user, err := createUserUseCase.Execute(username, email, password)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Email, email)
}

func TestCreateUserUseCase_FoundByEmail(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo)
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.UserEntity{}, nil)
	// act
	user, err := createUserUseCase.Execute(username, email, password)
	// assert
	assert.Nil(t, user)
	assert.Equal(t, err, exceptions.NewUserEmailAlreadyInUse())
}
