package usecases

import (
	"testing"
	"time"

	mock_providers "github.com/AndreyArthur/oganessone/src/application/providers/mocks"
	mock_repositories "github.com/AndreyArthur/oganessone/src/application/repositories/mocks"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
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

func TestCreateUserUseCase_SuccessCase(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := setup(t)
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
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(username, email, fakeBcryptHash).
		Return(repoUser, nil)
	repo.EXPECT().
		Save(repoUser).
		Return(nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user.IsValid())
}

func TestCreateUserUseCase_FoundByUsername(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
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
		FindByUsername(username, false).
		Return(nil, &shared.Error{})
	repo.EXPECT().
		FindByEmail(email).
		Return(&entities.UserEntity{}, nil)
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, user)
}

func TestCreateUserUseCase_FoundByEmail(t *testing.T) {
	// arrange
	useCase, repo, _, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password := "username", "user@email.com", "p4ssword"
	repo.EXPECT().
		FindByUsername(username, false).
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
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, &shared.Error{})
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, user)
}

func TestCreateUserUseCase_EncrypterHashReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := setup(t)
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
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, exceptions.NewInternalServerError())
	assert.Nil(t, user)
}

func TestCreateUserUseCase_CreateReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := setup(t)
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
		Create(username, email, fakeBcryptHash).
		Return(nil, &shared.Error{})
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Equal(t, err, &shared.Error{})
	assert.Nil(t, user)
}

func TestCreateUserUseCase_SaveReturnError(t *testing.T) {
	// arrange
	useCase, repo, encrypter, ctrl := setup(t)
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
		FindByUsername(username, false).
		Return(nil, nil)
	repo.EXPECT().
		FindByEmail(email).
		Return(nil, nil)
	encrypter.EXPECT().
		Hash(password).
		Return(fakeBcryptHash, nil)
	repo.EXPECT().
		Create(username, email, fakeBcryptHash).
		Return(repoUser, nil)
	repo.EXPECT().
		Save(repoUser).
		Return(&shared.Error{})
	// act
	user, err := useCase.Execute(username, email, password)
	// assert
	assert.Nil(t, user)
	assert.Equal(t, err, &shared.Error{})
}
