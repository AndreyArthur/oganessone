package presenters

import (
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type CreateSessionPresenter struct {
	createSession definitions.CreateSession
}

func (createSessionPresenter *CreateSessionPresenter) Handle(
	request *contracts.CreateSessionPresenterRequest,
) (*contracts.CreateSessionPresenterResponse, *shared.Error) {
	result, err := createSessionPresenter.createSession.
		Execute(&definitions.CreateSessionDTO{
			Login:    request.Body.Login,
			Password: request.Body.Password,
		})
	if err != nil {
		return nil, err
	}
	return &contracts.CreateSessionPresenterResponse{
		Body: &contracts.CreateSessionPresenterResponseBody{
			Account: &views.AccountView{
				Id:        result.Account.Id,
				Username:  result.Account.Username,
				Email:     result.Account.Email,
				CreatedAt: result.Account.CreatedAt.Format(time.RFC3339),
				UpdatedAt: result.Account.UpdatedAt.Format(time.RFC3339),
			},
			SessionKey: result.SessionKey,
		},
	}, nil
}

func NewCreateSessionPresenter(
	createSession definitions.CreateSession,
) (*CreateSessionPresenter, *shared.Error) {
	return &CreateSessionPresenter{
		createSession: createSession,
	}, nil
}
