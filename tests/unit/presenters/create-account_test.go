package test_presenters

import (
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	mock_definitions "github.com/AndreyArthur/oganessone/src/application/definitions/mocks"
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/presenters"
	"github.com/AndreyArthur/oganessone/tests/helpers/verifier"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type CreateUserPresenterTest struct{}

func (*CreateUserPresenterTest) setup(t *testing.T) (*presenters.CreateAccountPresenter, *mock_definitions.MockCreateAccount, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	useCase := mock_definitions.NewMockCreateAccount(ctrl)
	presenter, _ := presenters.NewCreateAccountPresenter(useCase)
	return presenter, useCase, ctrl
}

func TestCreateUserPresenter_SuccessCase(t *testing.T) {
	// arrange
	presenter, useCase, ctrl := (&CreateUserPresenterTest{}).setup(t)
	defer ctrl.Finish()
	uuid, _ := helpers.NewUuid()
	now := time.Now().UTC()
	id, username, email, password, createdAt, updatedAt :=
		uuid.Generate(),
		"username",
		"user@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW",
		now,
		now
	entity, _ := entities.NewAccountEntity(&dtos.AccountDTO{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	useCase.EXPECT().
		Execute(&definitions.CreateAccountDTO{
			Username: username,
			Email:    email,
			Password: password,
		}).
		Return(entity, nil)
	// act
	result, err := presenter.Handle(&contracts.CreateAccountPresenterRequest{
		Body: &contracts.CreateAccountPresenterRequestBody{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, err)
	assert.True(t, verifier.IsUuid(result.Body.Id))
	assert.True(t, verifier.IsAccountUsername(result.Body.Username))
	assert.True(t, verifier.IsEmail(result.Body.Email))
	assert.True(t, verifier.IsISO8601(result.Body.CreatedAt))
	assert.True(t, verifier.IsISO8601(result.Body.UpdatedAt))
}

func TestCreateUserPresenter_FailureCase(t *testing.T) {
	presenter, useCase, ctrl := (&CreateUserPresenterTest{}).setup(t)
	defer ctrl.Finish()
	username, email, password :=
		"username",
		"user@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	useCase.EXPECT().
		Execute(&definitions.CreateAccountDTO{
			Username: username,
			Email:    email,
			Password: password,
		}).
		Return(nil, &shared.Error{})
	// act
	result, err := presenter.Handle(&contracts.CreateAccountPresenterRequest{
		Body: &contracts.CreateAccountPresenterRequestBody{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}
