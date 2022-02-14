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

type CreateSessionPresenterTest struct{}

func (*CreateSessionPresenterTest) setup(t *testing.T) (*presenters.CreateSessionPresenter, *mock_definitions.MockCreateSession, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	useCase := mock_definitions.NewMockCreateSession(ctrl)
	presenter, _ := presenters.NewCreateSessionPresenter(useCase)
	return presenter, useCase, ctrl
}

func TestCreateSessionPresenter_SuccessCase(t *testing.T) {
	// arrange
	presenter, useCase, ctrl := (&CreateSessionPresenterTest{}).setup(t)
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
		Execute(&definitions.CreateSessionDTO{
			Login:    username,
			Password: password,
		}).
		Return(&definitions.CreateSessionResult{
			User:       entity,
			SessionKey: "session_key",
		}, nil)
	// act
	result, err := presenter.Handle(&contracts.CreateSessionPresenterRequest{
		Body: &contracts.CreateSessionPresenterRequestBody{
			Login:    username,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, err)
	assert.Equal(t, result.Body.SessionKey, "session_key")
	assert.True(t, verifier.IsUuid(result.Body.User.Id))
	assert.True(t, verifier.IsUserUsername(result.Body.User.Username))
	assert.True(t, verifier.IsEmail(result.Body.User.Email))
	assert.True(t, verifier.IsISO8601(result.Body.User.CreatedAt))
	assert.True(t, verifier.IsISO8601(result.Body.User.UpdatedAt))
}

func TestSessionUserPresenter_FailureCase(t *testing.T) {
	presenter, useCase, ctrl := (&CreateSessionPresenterTest{}).setup(t)
	defer ctrl.Finish()
	username, password :=
		"username",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	useCase.EXPECT().
		Execute(&definitions.CreateSessionDTO{
			Login:    username,
			Password: password,
		}).
		Return(nil, &shared.Error{})
	// act
	result, err := presenter.Handle(&contracts.CreateSessionPresenterRequest{
		Body: &contracts.CreateSessionPresenterRequestBody{
			Login:    username,
			Password: password,
		},
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}
