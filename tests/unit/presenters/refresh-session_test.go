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

type RefreshSessionPresenterTest struct{}

func (*RefreshSessionPresenterTest) setup(t *testing.T) (*presenters.RefreshSessionPresenter, *mock_definitions.MockRefreshSession, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	useCase := mock_definitions.NewMockRefreshSession(ctrl)
	presenter, _ := presenters.NewRefreshSessionPresenter(useCase)
	return presenter, useCase, ctrl
}

func TestRefreshSessionPresenter_SuccessCase(t *testing.T) {
	// arrange
	presenter, useCase, ctrl := (&RefreshSessionPresenterTest{}).setup(t)
	defer ctrl.Finish()
	uuid, _ := helpers.NewUuid()
	now := time.Now().UTC()
	sessionKey := "kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w"
	id, username, email, password, createdAt, updatedAt :=
		uuid.Generate(),
		"username",
		"account@email.com",
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
		Execute(&definitions.RefreshSessionDTO{
			SessionKey: sessionKey,
		}).
		Return(&definitions.RefreshSessionResult{
			Account:    entity,
			SessionKey: "session_key",
		}, nil)
	// act
	result, err := presenter.Handle(&contracts.RefreshSessionPresenterRequest{
		Body: &contracts.RefreshSessionPresenterRequestBody{
			SessionKey: sessionKey,
		},
	})
	// assert
	assert.Nil(t, err)
	assert.Equal(t, result.Body.SessionKey, "session_key")
	assert.True(t, verifier.IsUuid(result.Body.Account.Id))
	assert.True(t, verifier.IsAccountUsername(result.Body.Account.Username))
	assert.True(t, verifier.IsEmail(result.Body.Account.Email))
	assert.True(t, verifier.IsISO8601(result.Body.Account.CreatedAt))
	assert.True(t, verifier.IsISO8601(result.Body.Account.UpdatedAt))
}

func TestRefreshSessionPresenter_FailureCase(t *testing.T) {
	presenter, useCase, ctrl := (&RefreshSessionPresenterTest{}).setup(t)
	defer ctrl.Finish()
	sessionKey := "kAYWq$jq^Z^c/$D9$f~iLPD?*7F!w9(w"
	useCase.EXPECT().
		Execute(&definitions.RefreshSessionDTO{
			SessionKey: sessionKey,
		}).
		Return(nil, &shared.Error{})
	// act
	result, err := presenter.Handle(&contracts.RefreshSessionPresenterRequest{
		Body: &contracts.RefreshSessionPresenterRequestBody{
			SessionKey: sessionKey,
		},
	})
	// assert
	assert.Nil(t, result)
	assert.Equal(t, err, &shared.Error{})
}
