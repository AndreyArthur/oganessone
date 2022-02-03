package test_presenters

import (
	"regexp"
	"testing"
	"time"

	mock_definitions "github.com/AndreyArthur/oganessone/src/application/definitions/mocks"
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/presenters"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*presenters.CreateUserPresenter, *mock_definitions.MockCreateUser, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	useCase := mock_definitions.NewMockCreateUser(ctrl)
	presenter, _ := presenters.NewCreateUserPresenter(useCase)
	return presenter, useCase, ctrl
}

func TestCreateUserPresenter_SuccessCase(t *testing.T) {
	// arrange
	presenter, useCase, ctrl := setup(t)
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
	entity, _ := entities.NewUserEntity(&dtos.UserDTO{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	useCase.EXPECT().
		Execute(username, email, password).
		Return(entity, nil)
	iso8601Regex := regexp.MustCompile(`^\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(([+-]\d\d:\d\d)|Z)?$`)
	// act
	result, err := presenter.Handle(&contracts.CreateUserPresenterRequest{
		Body: &contracts.CreateUserPresenterRequestBody{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, err)
	assert.Equal(t, result.Body.Id, id)
	assert.Equal(t, result.Body.Username, username)
	assert.Equal(t, result.Body.Email, email)
	assert.True(t, iso8601Regex.Match([]byte(result.Body.CreatedAt)))
	assert.True(t, iso8601Regex.Match([]byte(result.Body.UpdatedAt)))
}

func TestCreateUserPresenter_FailCase(t *testing.T) {
	presenter, useCase, ctrl := setup(t)
	defer ctrl.Finish()
	username, email, password :=
		"username",
		"user@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	useCase.EXPECT().
		Execute(username, email, password).
		Return(nil, &shared.Error{})
	// act
	result, err := presenter.Handle(&contracts.CreateUserPresenterRequest{
		Body: &contracts.CreateUserPresenterRequestBody{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}
