package usecases

import (
	"testing"

	mock_repositories "github.com/AndreyArthur/murao-oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/murao-oganessone/src/core/entities"
	"github.com/AndreyArthur/murao-oganessone/src/core/exceptions"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_NotFoundByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo)
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(nil, nil)

	user, err := createUserUseCase.Execute(username, email, password)

	assert.Nil(t, err)
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Email, email)
}

func TestCreateUserUseCase_FoundByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repositories.NewMockUsersRepository(ctrl)
	createUserUseCase, _ := NewCreateUserUseCase(repo)
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, true).
		Return(&entities.UserEntity{}, nil)

	user, err := createUserUseCase.Execute(username, email, password)

	assert.Equal(t, err, exceptions.NewUserUsernameAlreadyInUse())
	assert.Nil(t, user)
}
